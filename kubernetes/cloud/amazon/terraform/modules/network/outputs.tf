output "vpc_id" {
  value = module.vpc.vpc_id
}

output "private_subnets" {
  value = module.vpc.private_subnets
}

output "public_subnets" {
  value = module.vpc.public_subnets
}

output "security_group_worker_1_id" {
  value = aws_security_group.node_ssh_group_1.id
}

output "security_group_worker_2_id" {
  value = aws_security_group.node_ssh_group_2.id
}

output "security_group_worker_all_id" {
  value = aws_security_group.node_ssh_all.id
}