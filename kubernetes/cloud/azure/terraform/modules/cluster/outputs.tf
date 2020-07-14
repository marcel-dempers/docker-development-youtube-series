output "kube_config" {
    value = azurerm_kubernetes_cluster.aks-getting-started.kube_config_raw
}

output "cluster_ca_certificate" {
    value = azurerm_kubernetes_cluster.aks-getting-started.kube_config.0.cluster_ca_certificate
}

output "client_certificate" {
    value = azurerm_kubernetes_cluster.aks-getting-started.kube_config.0.client_certificate
}

output "client_key" {
    value = azurerm_kubernetes_cluster.aks-getting-started.kube_config.0.client_key
}

output "host" {
    value = azurerm_kubernetes_cluster.aks-getting-started.kube_config.0.host
}