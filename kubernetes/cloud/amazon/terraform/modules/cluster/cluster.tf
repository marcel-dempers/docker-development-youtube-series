provider "local" {
  version = "~> 1.2"
}

provider "null" {
  version = "~> 2.1"
}

provider "template" {
  version = "~> 2.1"
}

module "eks" {
  source          = "terraform-aws-modules/eks/aws"
  version = "~> 12.1.0"
  cluster_name    = "eks-getting-started"
  cluster_version = "1.16"
  subnets         = var.private_subnets
  vpc_id          = var.vpc_id

  node_groups_defaults = {
    ami_type  = "AL2_x86_64"
    disk_size = 50
  }
  
  node_groups = {
    example = {
      desired_capacity = 1
      max_capacity     = 10
      min_capacity     = 1
      instance_type = "t2.small"
    }
  }
}


/*
  worker_additional_security_group_ids = [var.worker_group_all_security_id]
  worker_groups = [
    {
      name                          = "worker-group-1"
      instance_type                 = "t2.small"
      additional_userdata           = "echo foo bar"
      asg_desired_capacity          = 1
      additional_security_group_ids = [var.worker_group_1_security_id]
    },
  ]
}
*/

data "aws_eks_cluster" "cluster" {
  name = module.eks.cluster_id
}

data "aws_eks_cluster_auth" "cluster" {
  name = module.eks.cluster_id
}