terraform {
  required_providers {
    avi = {
      source  = "vmware/avi"
      version = "21.1.1"
    }
  }
}
provider "avi" {
  avi_username   = var.admin_username
  avi_password   = var.admin_password
  avi_controller = aws_instance.avi-controller[0].public_ip
  avi_tenant     = "admin"
  avi_version    = var.api_version
}
