#!/usr/bin/env bash
set -uo pipefail
export AWS_PAGER=""
REGION=${REGION:-eu-north-1}; PREFIX=tracing-demo; CLUSTER=$PREFIX-cluster
for s in order-service inventory-service; do aws ecs delete-service --region $REGION --cluster $CLUSTER --service $s --force >/dev/null; done
aws ecs wait services-inactive --region $REGION --cluster $CLUSTER --services order-service inventory-service
aws ecs delete-cluster --region $REGION --cluster $CLUSTER >/dev/null
NS_ID=$(aws servicediscovery list-namespaces --region $REGION --query "Namespaces[?Name=='$PREFIX.local'].Id" --output text)
for sd in $(aws servicediscovery list-services --region $REGION --filters Name=NAMESPACE_ID,Values=$NS_ID --query 'Services[].Id' --output text); do
  aws servicediscovery delete-service --region $REGION --id $sd; done
[ -n "$NS_ID" ] && aws servicediscovery delete-namespace --region $REGION --id $NS_ID >/dev/null
sleep 30
SG=$(aws ec2 describe-security-groups --region $REGION --filters Name=group-name,Values=$PREFIX-sg --query 'SecurityGroups[0].GroupId' --output text)
[ "$SG" != "None" ] && aws ec2 delete-security-group --region $REGION --group-id $SG
for r in inventory-service order-service; do aws ecr delete-repository --region $REGION --repository-name $PREFIX/$r --force >/dev/null; done
aws logs delete-log-group --region $REGION --log-group-name /ecs/$PREFIX
echo "teardown complete"
