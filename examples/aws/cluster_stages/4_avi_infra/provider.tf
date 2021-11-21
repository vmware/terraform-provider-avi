terraform {
  required_providers {
    avi = {
      source  = "vmware/avi"
      version = "21.1.1"
    }
  }
}
provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = data.aws_instance.avi_controller.public_ip
  avi_tenant     = "admin"
  avi_version    = var.api_version
}


