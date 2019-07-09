provider "azurerm" {
  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
}

data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
}

data "azurerm_network_interface" "controller_nic_2" {
  name                = "${var.project_name}-terraform-network-interface-2"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
}

data "azurerm_network_interface" "controller_nic_3" {
  name                = "${var.project_name}-terraform-network-interface-3"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
}

resource "avi_cluster" "azure_cluster" {
  name = "cluster-0-1"
  nodes {
    ip {
      type = "V4"
      addr = data.azurerm_network_interface.controller_nic.private_ip_address
    }
    name = "node01"
  }
  nodes {
    ip {
      type = "V4"
      addr = data.azurerm_network_interface.controller_nic_2.private_ip_address
    }
    name = "node02"
  }
  nodes {
    ip {
      type = "V4"
      addr = data.azurerm_network_interface.controller_nic_3.private_ip_address
    }
    name = "node03"
  }
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = data.azurerm_network_interface.controller_nic.private_ip_address
  avi_tenant     = "admin"
  avi_version    = var.avi_version
}

