output "host" {
  value = data.aws_eks_cluster.cluster.endpoint
}

output "token" {
  value = data.aws_eks_cluster_auth.cluster.token
}

output "cluster_ca_certificate" {
  value = base64decode(data.aws_eks_cluster.cluster.certificate_authority.0.data)
}