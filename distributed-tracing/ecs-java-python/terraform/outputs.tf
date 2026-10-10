output "cluster" {
  value = aws_ecs_cluster.this.name
}

output "images" {
  value = local.images
}

output "order_service_ip_command" {
  description = "Prints order-service's public IP (it changes whenever the task restarts)."
  value       = <<-EOT
    aws ecs describe-tasks --region ${var.region} --cluster ${aws_ecs_cluster.this.name} \
      --tasks $(aws ecs list-tasks --region ${var.region} --cluster ${aws_ecs_cluster.this.name} --service-name order-service --query 'taskArns[0]' --output text) \
      --query "tasks[0].attachments[0].details[?name=='networkInterfaceId'].value" --output text \
    | xargs -I{} aws ec2 describe-network-interfaces --region ${var.region} --network-interface-ids {} \
      --query 'NetworkInterfaces[0].Association.PublicIp' --output text
  EOT
}
