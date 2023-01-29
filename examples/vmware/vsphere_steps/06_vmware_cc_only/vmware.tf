data "vsphere_content_library" "library" {
  name =  var.content_library_name
}

resource "avi_cloud" "vmware_cloud_tf" {
  name         = var.cloud_name
  vtype        = "CLOUD_VCENTER"
  dhcp_enabled = true
  vcenter_configuration {
    use_content_lib = "true"
    username                 = var.vsphere_username
    content_lib {
          id = data.vsphere_content_library.library.id
        }
    datacenter               = var.vcenter_datacenter
    management_network       = var.vcenter_network
    privilege                = var.vsphere_privilege
    vcenter_url              = var.vcenter_url
    password                 = var.vsphere_password
    deactivate_vm_discovery  =  "false"
  }
  license_tier = var.avi_license
  license_type = var.vcenter_license_type
  tenant_ref   = var.tenant
  maintenance_mode           = "false"
}
