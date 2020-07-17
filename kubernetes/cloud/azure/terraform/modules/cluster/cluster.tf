resource "azurerm_resource_group" "aks-getting-started" {
  name     = "aks-getting-started"
  location = var.location
}

resource "azurerm_kubernetes_cluster" "aks-getting-started" {
  name                  = "aks-getting-started"
  location              = azurerm_resource_group.aks-getting-started.location
  resource_group_name   = azurerm_resource_group.aks-getting-started.name
  dns_prefix            = "aks-getting-started"            
  kubernetes_version    =  var.kubernetes_version
  
  default_node_pool {
    name       = "default"
    node_count = 1
    vm_size    = "Standard_E4s_v3"
    type       = "VirtualMachineScaleSets"
    os_disk_size_gb = 250
  }

  service_principal  {
    client_id = var.serviceprinciple_id
    client_secret = var.serviceprinciple_key
  }

  linux_profile {
    admin_username = "azureuser"
    ssh_key {
        key_data = var.ssh_key
    }
  }

  network_profile {
      network_plugin = "kubenet"
      load_balancer_sku = "Standard"
  }

  addon_profile {
    aci_connector_linux {
      enabled = false
    }

    azure_policy {
      enabled = false
    }

    http_application_routing {
      enabled = false
    }

    kube_dashboard {
      enabled = false
    }

    oms_agent {
      enabled = false
    }
  }

}

/*
resource "azurerm_kubernetes_cluster_node_pool" "monitoring" {
  name                  = "monitoring"
  kubernetes_cluster_id = azurerm_kubernetes_cluster.aks-getting-started.id
  vm_size               = "Standard_DS2_v2"
  node_count            = 1
  os_disk_size_gb       = 250
  os_type               = "Linux"
}

*/