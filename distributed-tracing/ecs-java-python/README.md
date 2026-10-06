# ECS distributed tracing demo (Java → Python)

Two services on AWS ECS Fargate, auto-instrumented with Middleware via the
[`mw-ecs`](https://github.com/middleware-labs/mw-ecs-instrumentation) CLI. Built to demo
end-to-end distributed tracing on ECS and OpsAI fixing a code-level bug.

```
order-service (Java 21 / Spring Boot, :8080)
  GET /api/orders/quote/{productId}
    ├─► GET http://inventory.tracing-demo.local:5000/api/products/{id}          → 200
    └─► GET http://inventory.tracing-demo.local:5000/api/products/{id}/rating   → 500 for P400 / P500
inventory-service (Python 3.12 / Flask + gunicorn, :5000)
```

- **Intentional bug:** `average_rating()` in `inventory-service-python/app.py` divides by
  `len(reviews)`. P400 and P500 have no reviews, so it raises `ZeroDivisionError` → HTTP 500,
  and order-service returns 502. Roughly 40% of requests fail.
- **Traffic:** order-service calls its own endpoint every 3s with a random product
  (`TRAFFIC_ENABLED`, `TRAFFIC_INTERVAL_MS`), so each trace starts at an inbound Java request.
- **Service discovery:** Cloud Map private DNS namespace `tracing-demo.local`.
- The app code has no instrumentation; it is all added by `mw-ecs`.

## Prerequisites

- AWS CLI with credentials, Docker, Python 3
- Default VPC in the target region and the `ecsTaskExecutionRole` IAM role
- Middleware API key and target URL (Installation → Agent → AWS ECS → Auto-Instrumentation)

## Run

```bash
# 1. Build, push to ECR, and create the cluster + both services (Fargate ARM64)
REGION=eu-north-1 ./deploy.sh

# 2. Auto-instrument both services with Middleware
MW_API_KEY=<key> MW_TARGET=https://<uid>.middleware.io:443 ./instrument.sh

# 3. Hit it manually (deploy.sh prints the order-service IP; port 8080 is open to your IP only)
curl http://<ip>:8080/api/orders/quote/P100   # 200
curl http://<ip>:8080/api/orders/quote/P400   # 502, ZeroDivisionError in inventory-service
```

In Middleware: **APM → Services → order-service → Service Map** shows
`order-service → inventory-service`; open any `ZeroDivisionError` span and use the **Errors**
tab to see the failing function body and **Fix with OpsAI**.

After fixing the bug, run `./deploy.sh` again. On existing services it only rolls out the new
`:latest` image and keeps the instrumented task definition.

`./teardown.sh` removes everything (services, cluster, Cloud Map namespace, security group,
ECR repos, log group).

## What `instrument.sh` does

Runs `mw-ecs instrument --enable-apm --enable-logs` for each service, which adds an `mw-agent`
sidecar, an OTel auto-instrumentation init container (Java / Python), and a FireLens
`log_router`. On top of that it:

| Change | Why |
|---|---|
| Restores `runtimePlatform: ARM64` | mw-ecs v0.0.1 drops it, so Fargate would default to x86 and the arm64 images would fail to start |
| Answers `y` to "Override with awsfirelens?" | With an existing CloudWatch log config, the default `N` produces a task definition ECS rejects (FireLens router with no `awsfirelens` container) |
| Uses `autoinstrumentation-python:0.64b0-opsai` and sets `MW_RECORD_EXCEPTION_SOURCE=true` | Captures the function body on exceptions (off by default in that image) |
| Adds CloudWatch logs to the sidecars | So `mw-agent` / `log_router` / init container can be debugged |

## AWS resources created

All prefixed `tracing-demo`: ECS cluster `tracing-demo-cluster`, ECR repos
`tracing-demo/{inventory,order}-service`, Cloud Map namespace `tracing-demo.local`,
security group `tracing-demo-sg` (5000 from itself, 8080 from your IP), log group
`/ecs/tracing-demo`.
