locals {
  mw_volume = "mw-agent-instrumentation"

  # Per-service APM wiring for the OTel auto-instrumentation init container.
  apm = {
    inventory-service = {
      init_image = var.python_autoinstrumentation_image
      mount      = "/otel-auto-instrumentation-python"
      init_cmd   = ["cp", "-r", "/autoinstrumentation/.", "/otel-auto-instrumentation-python"]
      env = {
        PYTHONPATH                 = "/otel-auto-instrumentation-python/opentelemetry/instrumentation/auto_instrumentation:/otel-auto-instrumentation-python"
        MW_RECORD_EXCEPTION_SOURCE = "true"
      }
    }
    order-service = {
      init_image = var.java_autoinstrumentation_image
      mount      = "/otel-auto-instrumentation-java"
      init_cmd   = ["cp", "/javaagent.jar", "/otel-auto-instrumentation-java/javaagent.jar"]
      env        = { JAVA_TOOL_OPTIONS = "-javaagent:/otel-auto-instrumentation-java/javaagent.jar" }
    }
  }

  app_env = {
    inventory-service = {}
    order-service = {
      INVENTORY_URL       = "http://inventory.${local.namespace}:5000"
      TRAFFIC_INTERVAL_MS = tostring(var.traffic_interval_ms)
    }
  }

  awslogs = { for name, _ in local.services : name => {
    logDriver = "awslogs"
    options = {
      "awslogs-group"         = aws_cloudwatch_log_group.this.name
      "awslogs-region"        = var.region
      "awslogs-stream-prefix" = name
    }
  } }

  containers = { for name, svc in local.services : name => concat(
    [{
      name         = name
      image        = local.images[name]
      essential    = true
      portMappings = [{ containerPort = svc.port, protocol = "tcp" }]
      environment = [for k, v in merge(
        local.app_env[name],
        var.enable_middleware ? merge(local.apm[name].env, {
          OTEL_SERVICE_NAME           = name
          OTEL_EXPORTER_OTLP_ENDPOINT = "http://localhost:9320"
          OTEL_EXPORTER_OTLP_PROTOCOL = "http/protobuf"
          OTEL_RESOURCE_ATTRIBUTES    = "mw.account_key=${var.mw_api_key}"
        }) : {}
      ) : { name = k, value = v }]
      mountPoints = jsondecode(var.enable_middleware ? jsonencode([{ sourceVolume = local.mw_volume, containerPath = local.apm[name].mount, readOnly = true }]) : "[]")
      dependsOn   = jsondecode(var.enable_middleware ? jsonencode([{ containerName = "instrumentation-init", condition = "SUCCESS" }]) : "[]")
      # With Middleware, app logs go to the mw-agent via FireLens; otherwise to CloudWatch.
      logConfiguration = jsondecode(var.enable_middleware ? jsonencode({
        logDriver = "awsfirelens"
        options   = { Name = "forward", Host = "127.0.0.1", Port = "8006" }
      }) : jsonencode(local.awslogs[name]))
    }],
    # jsonencode/jsondecode lets both branches share one type (HCL rejects tuples of different shapes).
    jsondecode(var.enable_middleware ? jsonencode([
      {
        name      = "mw-agent"
        image     = var.mw_agent_image
        cpu       = 256
        essential = false
        portMappings = [
          { containerPort = 8006, protocol = "tcp", appProtocol = "http" },
          { containerPort = 9320, protocol = "tcp" },
        ]
        environment = [
          { name = "MW_API_KEY", value = var.mw_api_key },
          { name = "MW_TARGET", value = var.mw_target },
          { name = "OTEL_EXPORTER_OTLP_PROTOCOL", value = "grpc" },
        ]
        logConfiguration = merge(local.awslogs[name], { options = merge(local.awslogs[name].options, { "awslogs-stream-prefix" = "${name}-mw-agent" }) })
      },
      {
        name                  = "log_router"
        image                 = "public.ecr.aws/aws-observability/aws-for-fluent-bit:stable"
        essential             = true
        user                  = "0"
        firelensConfiguration = { type = "fluentbit", options = { "enable-ecs-log-metadata" = "true" } }
        logConfiguration      = merge(local.awslogs[name], { options = merge(local.awslogs[name].options, { "awslogs-stream-prefix" = "${name}-log_router" }) })
      },
      {
        name             = "instrumentation-init"
        image            = local.apm[name].init_image
        cpu              = 128
        memory           = 128
        essential        = false
        command          = local.apm[name].init_cmd
        mountPoints      = [{ sourceVolume = local.mw_volume, containerPath = local.apm[name].mount, readOnly = false }]
        logConfiguration = merge(local.awslogs[name], { options = merge(local.awslogs[name].options, { "awslogs-stream-prefix" = "${name}-instrumentation-init" }) })
      },
    ]) : "[]")
  ) }
}

resource "aws_ecs_task_definition" "app" {
  for_each                 = local.services
  family                   = "${var.prefix}-${each.key}"
  network_mode             = "awsvpc"
  requires_compatibilities = ["FARGATE"]
  cpu                      = "512"
  memory                   = "1024"
  execution_role_arn       = aws_iam_role.execution.arn
  container_definitions    = jsonencode(local.containers[each.key])

  runtime_platform {
    cpu_architecture        = "ARM64"
    operating_system_family = "LINUX"
  }

  dynamic "volume" {
    for_each = var.enable_middleware ? [1] : []
    content {
      name = local.mw_volume
    }
  }

  lifecycle {
    precondition {
      condition     = !var.enable_middleware || var.mw_api_key != ""
      error_message = "mw_api_key is required when enable_middleware = true."
    }
  }

  depends_on = [null_resource.image]
}

resource "aws_ecs_service" "app" {
  for_each        = local.services
  name            = each.key
  cluster         = aws_ecs_cluster.this.id
  task_definition = aws_ecs_task_definition.app[each.key].arn
  desired_count   = 1
  launch_type     = "FARGATE"

  network_configuration {
    subnets          = data.aws_subnets.default.ids
    security_groups  = [aws_security_group.this.id]
    assign_public_ip = true
  }

  dynamic "service_registries" {
    for_each = each.key == "inventory-service" ? [1] : []
    content {
      registry_arn = aws_service_discovery_service.inventory.arn
    }
  }

  wait_for_steady_state = true
}
