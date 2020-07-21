provider "aws" {
  version = ">= 2.28.1"
  region     = var.region
  access_key = var.access_key
  secret_key = var.secret_key
}

module "network" {
  source                      = "./modules/network/"
}


module "cluster" {
  source                      = "./modules/cluster/"
  vpc_id                      = "${module.network.vpc_id}"
  private_subnets             = "${module.network.private_subnets}"
  public_subnets              = "${module.network.public_subnets}"
  #worker_group_1_security_id  = "${module.network.security_group_worker_1_id}"
  #worker_group_all_security_id= "${module.network.security_group_worker_all_id}"

  #location              = var.location
  #kubernetes_version    = var.kubernetes_version  
  
}

module "k8s" {
  source                = "./modules/k8s/"
  host                  = "${module.cluster.host}"
  token                 = "${module.cluster.token}"
  cluster_ca_certificate= "${module.cluster.cluster_ca_certificate}"
}
