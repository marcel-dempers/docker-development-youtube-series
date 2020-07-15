provider "azurerm" {
  version = "=2.0.0"
  features {}
}

module "cluster" {
  source                = "./modules/cluster/"
  serviceprinciple_id   = var.serviceprinciple_id
  serviceprinciple_key  = var.serviceprinciple_key
  ssh_key               = var.ssh_key
  location              = var.location
  kubernetes_version    = var.kubernetes_version  
  
}

module "k8s" {
  source                = "./modules/k8s/"
  host                  = "${module.cluster.host}"
  client_certificate    = "${base64decode(module.cluster.client_certificate)}"
  client_key            = "${base64decode(module.cluster.client_key)}"
  cluster_ca_certificate= "${base64decode(module.cluster.cluster_ca_certificate)}"

}

module "k8s_monitoring_prometheus_operator" {
  source                = "./modules/monitoring/prometheus-operator/"
  host                  = "${module.cluster.host}"
  client_certificate    = "${base64decode(module.cluster.client_certificate)}"
  client_key            = "${base64decode(module.cluster.client_key)}"
  cluster_ca_certificate= "${base64decode(module.cluster.cluster_ca_certificate)}"

}