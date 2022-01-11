terraform {
  required_providers {
    avi = {
      source  = "vmware/avi"
      version = "21.1.1"
    }
  }
}
provider "avi" {
  avi_username    = var.avi_username
  avi_tenant      = var.avi_tenant
  avi_password    = var.avi_password
  avi_controller  = var.avi_controller
  avi_version     = var.avi_version
  avi_api_timeout = 50
}

