terraform {
  required_providers {
    avi = {
      source  = "vmware/avi"
      version = "21.1.3"
    }
  }
}
provider "aws" {
  shared_credentials_files = [var.aws_creds_file]
  region                  = var.aws_region
}
provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_current_password
  avi_controller = data.aws_instance.avi_controller[0].public_ip
  avi_tenant     = var.avi_tenant
}

