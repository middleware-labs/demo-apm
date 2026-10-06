#!/usr/bin/env bash
# Auto-instruments both ECS services with Middleware using the mw-ecs CLI
# (https://github.com/middleware-labs/mw-ecs-instrumentation), then rolls the services.
#
#   MW_API_KEY=<key> MW_TARGET=https://<uid>.middleware.io:443 ./instrument.sh
#
# Workarounds applied on top of the mw-ecs output (v0.0.1):
#   - restores runtimePlatform (ARM64), which mw-ecs drops from the generated task definition
#   - answers "y" to "Override with awsfirelens?" (the default "N" yields an unregistrable task definition)
#   - swaps the Python init container for the OpsAI build and enables exception source capture
set -euo pipefail
export AWS_PAGER=""
: "${MW_API_KEY:?set MW_API_KEY}"
: "${MW_TARGET:?set MW_TARGET}"
REGION=${REGION:-eu-north-1}
PREFIX=tracing-demo
CLUSTER=$PREFIX-cluster
PYTHON_INIT_IMAGE=${PYTHON_INIT_IMAGE:-ghcr.io/middleware-labs/opentelemetry-operator/autoinstrumentation-python:0.64b0-opsai}
cd "$(dirname "$0")"
WORK=$(mktemp -d); trap 'rm -rf "$WORK"' EXIT

MW_ECS=$(command -v mw-ecs || true)
if [ -z "$MW_ECS" ]; then
  echo "==> Downloading mw-ecs CLI"
  OS=$(uname -s | tr '[:upper:]' '[:lower:]')
  ARCH=$(uname -m); case "$ARCH" in x86_64|amd64) ARCH=amd64 ;; aarch64|arm64) ARCH=arm64 ;; esac
  VERSION=$(curl -fsSL https://api.github.com/repos/middleware-labs/mw-ecs-instrumentation/releases/latest | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/')
  MW_ECS=$WORK/mw-ecs
  curl -fsSL -o "$MW_ECS" "https://github.com/middleware-labs/mw-ecs-instrumentation/releases/download/$VERSION/mw-ecs-$OS-$ARCH"
  chmod +x "$MW_ECS"
fi

instrument() { # service language
  local svc=$1 lang=$2 current out
  current=$(aws ecs describe-services --region $REGION --cluster $CLUSTER --services $svc --query 'services[0].taskDefinition' --output text)
  if aws ecs describe-task-definition --region $REGION --task-definition "$current" \
       --query "taskDefinition.containerDefinitions[?name=='mw-agent'].name" --output text | grep -q mw-agent; then
    echo "==> $svc already instrumented ($current), skipping"; return
  fi
  echo "==> Instrumenting $svc ($current, $lang)"
  out=$WORK/$svc.json
  printf 'y\n' | "$MW_ECS" instrument --region $REGION --fargate --enable-apm --enable-logs \
    --task-definition "$current" --language $lang --libc glibc --service-name $svc \
    --mw-api-key "$MW_API_KEY" --mw-target "$MW_TARGET" --output "$out" >"$WORK/$svc.log" 2>&1 \
    || { sed "s/$MW_API_KEY/***/g" "$WORK/$svc.log"; exit 1; }

  python3 - "$out" "$svc" "$REGION" "$PYTHON_INIT_IMAGE" <<'EOF'
import json, sys
path, svc, region, py_image = sys.argv[1:]
td = json.load(open(path))
td["runtimePlatform"] = {"cpuArchitecture": "ARM64", "operatingSystemFamily": "LINUX"}
for c in td["containerDefinitions"]:
    if c["name"] == svc:
        continue
    # Sidecars get CloudWatch logs so they can be debugged; the app container logs go to Middleware via FireLens.
    c.setdefault("logConfiguration", {"logDriver": "awslogs", "options": {
        "awslogs-group": "/ecs/tracing-demo", "awslogs-region": region, "awslogs-stream-prefix": f"{svc}-{c['name']}"}})
    if c["name"] == "instrumentation-init" and "autoinstrumentation-python" in c["image"]:
        c["image"] = py_image
        app = next(a for a in td["containerDefinitions"] if a["name"] == svc)
        app["environment"] = [e for e in app.get("environment", []) if e["name"] != "MW_RECORD_EXCEPTION_SOURCE"]
        app["environment"].append({"name": "MW_RECORD_EXCEPTION_SOURCE", "value": "true"})
json.dump(td, open(path, "w"), indent=2)
EOF

  local rev
  rev=$(aws ecs register-task-definition --region $REGION --cli-input-json "file://$out" --query taskDefinition.revision --output text)
  aws ecs update-service --region $REGION --cluster $CLUSTER --service $svc \
    --task-definition $PREFIX-$svc:$rev --force-new-deployment >/dev/null
  echo "    registered and deployed $PREFIX-$svc:$rev"
}

instrument inventory-service python
instrument order-service java

echo "==> Waiting for services to stabilize"
aws ecs wait services-stable --region $REGION --cluster $CLUSTER --services inventory-service order-service
echo "done - traces should appear in Middleware APM within a minute"
