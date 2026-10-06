#!/usr/bin/env bash
# Deploys order-service (Java) -> inventory-service (Python) to ECS Fargate (ARM64).
set -euo pipefail
export AWS_PAGER=""
REGION=${REGION:-eu-north-1}
PREFIX=tracing-demo
CLUSTER=$PREFIX-cluster
NAMESPACE=$PREFIX.local
ACCOUNT=$(aws sts get-caller-identity --query Account --output text)
REGISTRY=$ACCOUNT.dkr.ecr.$REGION.amazonaws.com
EXEC_ROLE=$(aws iam get-role --role-name ecsTaskExecutionRole --query Role.Arn --output text)
VPC=$(aws ec2 describe-vpcs --region $REGION --filters Name=isDefault,Values=true --query 'Vpcs[0].VpcId' --output text)
SUBNETS=$(aws ec2 describe-subnets --region $REGION --filters Name=default-for-az,Values=true --query 'Subnets[].SubnetId' --output text | tr '\t' ',')
MY_IP=$(curl -s https://checkip.amazonaws.com)
cd "$(dirname "$0")"

echo "==> ECR repos + images"
aws ecr get-login-password --region $REGION | docker login -u AWS --password-stdin $REGISTRY >/dev/null
for svc in inventory-service:inventory-service-python order-service:order-service-java; do
  name=${svc%%:*}; dir=${svc##*:}
  aws ecr describe-repositories --region $REGION --repository-names $PREFIX/$name >/dev/null 2>&1 \
    || aws ecr create-repository --region $REGION --repository-name $PREFIX/$name >/dev/null
  docker build --platform linux/arm64 -q -t $REGISTRY/$PREFIX/$name:latest ./$dir
  docker push -q $REGISTRY/$PREFIX/$name:latest
done

echo "==> Cluster, logs, security group"
aws ecs create-cluster --region $REGION --cluster-name $CLUSTER >/dev/null
aws logs create-log-group --region $REGION --log-group-name /ecs/$PREFIX 2>/dev/null || true
SG=$(aws ec2 describe-security-groups --region $REGION --filters Name=group-name,Values=$PREFIX-sg Name=vpc-id,Values=$VPC --query 'SecurityGroups[0].GroupId' --output text)
if [ "$SG" = "None" ]; then
  SG=$(aws ec2 create-security-group --region $REGION --group-name $PREFIX-sg --description "ECS tracing demo" --vpc-id $VPC --query GroupId --output text)
  aws ec2 authorize-security-group-ingress --region $REGION --group-id $SG --protocol tcp --port 5000 --source-group $SG >/dev/null
  aws ec2 authorize-security-group-ingress --region $REGION --group-id $SG --protocol tcp --port 8080 --cidr $MY_IP/32 >/dev/null
fi

echo "==> Cloud Map namespace $NAMESPACE"
NS_ID=$(aws servicediscovery list-namespaces --region $REGION --query "Namespaces[?Name=='$NAMESPACE'].Id" --output text)
if [ -z "$NS_ID" ]; then
  OP=$(aws servicediscovery create-private-dns-namespace --region $REGION --name $NAMESPACE --vpc $VPC --query OperationId --output text)
  until [ "$(aws servicediscovery get-operation --region $REGION --operation-id $OP --query Operation.Status --output text)" = "SUCCESS" ]; do sleep 5; done
  NS_ID=$(aws servicediscovery list-namespaces --region $REGION --query "Namespaces[?Name=='$NAMESPACE'].Id" --output text)
fi
SD_ARN=$(aws servicediscovery list-services --region $REGION --filters Name=NAMESPACE_ID,Values=$NS_ID --query "Services[?Name=='inventory'].Arn" --output text)
if [ -z "$SD_ARN" ]; then
  SD_ARN=$(aws servicediscovery create-service --region $REGION --name inventory --namespace-id $NS_ID \
    --dns-config 'RoutingPolicy=MULTIVALUE,DnsRecords=[{Type=A,TTL=10}]' \
    --health-check-custom-config FailureThreshold=1 --query Service.Arn --output text)
fi

echo "==> Task definitions (only used when a service is first created)"
taskdef() { # name port image env-json
  cat <<JSON
{
  "family": "$PREFIX-$1",
  "networkMode": "awsvpc",
  "requiresCompatibilities": ["FARGATE"],
  "runtimePlatform": {"cpuArchitecture": "ARM64", "operatingSystemFamily": "LINUX"},
  "cpu": "512", "memory": "1024",
  "executionRoleArn": "$EXEC_ROLE",
  "containerDefinitions": [{
    "name": "$1",
    "image": "$REGISTRY/$PREFIX/$1:latest",
    "essential": true,
    "portMappings": [{"containerPort": $2, "protocol": "tcp"}],
    "environment": $3,
    "logConfiguration": {"logDriver": "awslogs", "options": {
      "awslogs-group": "/ecs/$PREFIX", "awslogs-region": "$REGION", "awslogs-stream-prefix": "$1"}}
  }]
}
JSON
}
aws ecs register-task-definition --region $REGION --cli-input-json "$(taskdef inventory-service 5000 '[]')" >/dev/null
ORDER_ENV='[{"name":"INVENTORY_URL","value":"http://inventory.'$NAMESPACE':5000"}]'
aws ecs register-task-definition --region $REGION --cli-input-json "$(taskdef order-service 8080 "$ORDER_ENV")" >/dev/null

echo "==> Services"
NET="awsvpcConfiguration={subnets=[$SUBNETS],securityGroups=[$SG],assignPublicIp=ENABLED}"
create_or_update() { # name [extra args...]
  local name=$1; shift
  if [ "$(aws ecs describe-services --region $REGION --cluster $CLUSTER --services $name --query 'services[0].status' --output text)" = "ACTIVE" ]; then
    # Keep the service's current (Middleware-instrumented) task definition; just pull the new :latest image.
    aws ecs update-service --region $REGION --cluster $CLUSTER --service $name --force-new-deployment >/dev/null
  else
    aws ecs create-service --region $REGION --cluster $CLUSTER --service-name $name --task-definition $PREFIX-$name \
      --desired-count 1 --launch-type FARGATE --network-configuration "$NET" "$@" >/dev/null
  fi
}
create_or_update inventory-service --service-registries registryArn=$SD_ARN
create_or_update order-service

echo "==> Waiting for services to stabilize"
aws ecs wait services-stable --region $REGION --cluster $CLUSTER --services inventory-service order-service
TASK=$(aws ecs list-tasks --region $REGION --cluster $CLUSTER --service-name order-service --query 'taskArns[0]' --output text)
ENI=$(aws ecs describe-tasks --region $REGION --cluster $CLUSTER --tasks $TASK --query "tasks[0].attachments[0].details[?name=='networkInterfaceId'].value" --output text)
PUB=$(aws ec2 describe-network-interfaces --region $REGION --network-interface-ids $ENI --query 'NetworkInterfaces[0].Association.PublicIp' --output text)
echo "order-service: http://$PUB:8080/api/orders/quote/P100  (P400/P500 hit the bug)"
