terraform {
  required_providers {
    avi = {
      source  = "vmware/avi"
      version = "21.1.1"
    }
  }
}
provider "azurerm" {
  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
  features {}
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = data.azurerm_public_ip.avi_pip.ip_address
  avi_tenant     = "admin"
  avi_version    = var.avi_version
}

