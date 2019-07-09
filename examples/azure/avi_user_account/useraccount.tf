provider "azurerm" {
  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
}

data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface"
  resource_group_name = "${var.project_name}-terraform-resource-group"
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_current_password
  avi_controller = data.azurerm_network_interface.controller_nic.private_ip_address
  avi_tenant     = "admin"
  avi_version    = var.avi_version
}

resource "avi_useraccount" "avi_user" {
  username     = var.avi_username
  old_password = var.avi_current_password
  password     = var.avi_new_password
}

