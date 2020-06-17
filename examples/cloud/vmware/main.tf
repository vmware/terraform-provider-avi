provider "avi" {                                                                                                                                                        
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = var.avi_controller
  avi_tenant     = var.tenant
  avi_version    = var.avi_version
}

data "avi_tenant" "tenant" {
  name = var.tenant
}

resource "avi_cloud" "vmware_cloud_tf_1" {
  name         = var.cloud_name
  vtype        = "CLOUD_VCENTER"
  dhcp_enabled = true
  vcenter_configuration {
    username = var.vcenter_username
    datacenter = var.vcenter_datacenter
    management_network = var.vcenter_management_network
    privilege = var.vcenter_privilege
    vcenter_url = var.vcenter_vcenter_url
    password = var.vcenter_password
  }
  license_tier = var.vcenter_license_tier
  license_type = var.vcenter_license_type
  tenant_ref   = data.avi_tenant.tenant.id
}
