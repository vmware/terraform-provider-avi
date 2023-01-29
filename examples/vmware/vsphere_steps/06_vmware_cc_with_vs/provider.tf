terraform {
  required_providers {
    avi = {
      source  = "vmware/avi"
      version = "22.1.1-p1.0"
    }
  }
}
provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller  = var.avi_controller_ips[0]
  avi_tenant     = var.tenant
  avi_version    = var.avi_version
  avi_api_timeout = 50
}
provider "vsphere" {
  vsphere_server = var.vsphere_url
  user           = var.vsphere_username
  password       = var.vsphere_password
  allow_unverified_ssl = true
}
