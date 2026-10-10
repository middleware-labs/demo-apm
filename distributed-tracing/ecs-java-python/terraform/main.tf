data "aws_caller_identity" "current" {}

data "aws_vpc" "default" {
  default = true
}

data "aws_subnets" "default" {
  filter {
    name   = "vpc-id"
    values = [data.aws_vpc.default.id]
  }
  filter {
    name   = "default-for-az"
    values = ["true"]
  }
}

data "http" "my_ip" {
  count = var.allowed_cidr == null ? 1 : 0
  url   = "https://checkip.amazonaws.com"
}

locals {
  registry     = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${var.region}.amazonaws.com"
  allowed_cidr = coalesce(var.allowed_cidr, "${try(trimspace(data.http.my_ip[0].response_body), "")}/32")
  namespace    = "${var.prefix}.local"

  services = {
    inventory-service = { dir = "inventory-service-python", port = 5000 }
    order-service     = { dir = "order-service-java", port = 8080 }
  }

  # Image tag = hash of the service's source, so code changes roll a new task definition.
  image_tags = {
    for name, svc in local.services : name => substr(sha1(join("", [
      for f in sort(fileset("${path.module}/../${svc.dir}", "**")) : filesha1("${path.module}/../${svc.dir}/${f}")
    ])), 0, 12)
  }
  images = { for name, _ in local.services : name => "${aws_ecr_repository.app[name].repository_url}:${local.image_tags[name]}" }
}

# ── ECR + image build ─────────────────────────────────────────
resource "aws_ecr_repository" "app" {
  for_each     = local.services
  name         = "${var.prefix}/${each.key}"
  force_delete = true
}

resource "null_resource" "image" {
  for_each = local.services
  triggers = { image = local.images[each.key] }

  provisioner "local-exec" {
    command = <<-EOT
      set -e
      aws ecr get-login-password --region ${var.region} | docker login -u AWS --password-stdin ${local.registry} >/dev/null
      docker build --platform linux/arm64 -t ${local.images[each.key]} ${path.module}/../${each.value.dir}
      docker push ${local.images[each.key]}
    EOT
  }
}

# ── Cluster, IAM, logs ────────────────────────────────────────
resource "aws_ecs_cluster" "this" {
  name = "${var.prefix}-cluster"
}

resource "aws_cloudwatch_log_group" "this" {
  name              = "/ecs/${var.prefix}"
  retention_in_days = 7
}

resource "aws_iam_role" "execution" {
  name = "${var.prefix}-task-execution"
  assume_role_policy = jsonencode({
    Version   = "2012-10-17"
    Statement = [{ Effect = "Allow", Action = "sts:AssumeRole", Principal = { Service = "ecs-tasks.amazonaws.com" } }]
  })
}

resource "aws_iam_role_policy_attachment" "execution" {
  role       = aws_iam_role.execution.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonECSTaskExecutionRolePolicy"
}

# ── Networking ────────────────────────────────────────────────
resource "aws_security_group" "this" {
  name        = "${var.prefix}-sg"
  description = "ECS tracing demo"
  vpc_id      = data.aws_vpc.default.id

  ingress {
    description = "inventory-service from tasks in this SG"
    from_port   = 5000
    to_port     = 5000
    protocol    = "tcp"
    self        = true
  }
  ingress {
    description = "order-service from allowed CIDR"
    from_port   = 8080
    to_port     = 8080
    protocol    = "tcp"
    cidr_blocks = [local.allowed_cidr]
  }
  egress {
    from_port   = 0
    to_port     = 0
    protocol    = "-1"
    cidr_blocks = ["0.0.0.0/0"]
  }
}

resource "aws_service_discovery_private_dns_namespace" "this" {
  name = local.namespace
  vpc  = data.aws_vpc.default.id
}

resource "aws_service_discovery_service" "inventory" {
  name = "inventory"
  dns_config {
    namespace_id   = aws_service_discovery_private_dns_namespace.this.id
    routing_policy = "MULTIVALUE"
    dns_records {
      type = "A"
      ttl  = 10
    }
  }
  health_check_custom_config {
    failure_threshold = 1
  }
}
