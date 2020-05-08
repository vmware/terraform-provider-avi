provider "vsphere" {
  user           = var.vsphere_user
  password       = var.vsphere_password
  vsphere_server = var.vsphere_server
  allow_unverified_ssl = true
  version = "1.15.0"
}

module "vmware_deploy" {
  source = "../../../modules/services/vmware_deploy"

  vm_datastore = var.vm_datastore
  vm_resource_pool = var.vm_resource_pool
  vm_datacenter = var.vm_datacenter
  vm_network = var.vm_network
  vm_template = var.vm_template
  vm_name = var.vm_name
  vm_folder = var.vm_folder

  avi_new_password = var.avi_new_password
  avi_password = var.avi_password
  avi_username = var.avi_username
}

output "controller_1" {
  value = module.vmware_deploy.vsphere_virtual_machine_vm1
}

output "controller_2" {
  value = module.vmware_deploy.vsphere_virtual_machine_vm2
}

output "controller_3" {
  value = module.vmware_deploy.vsphere_virtual_machine_vm3
}

output "avi_cluster_output" {
  value = module.vmware_deploy.avi_cluster_output
}

