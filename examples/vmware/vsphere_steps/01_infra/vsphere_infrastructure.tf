
data "vsphere_datacenter" "dc" {
  name = var.vcenter_dc
}

# data "vsphere_compute_cluster" "compute_cluster" {
#   name          = var.vcenter_cluster
#   datacenter_id = data.vsphere_datacenter.dc.id
# }

data "vsphere_datastore" "datastore" {
  name = var.vcenter_datastore
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_compute_cluster" "cluster" {
   name          = var.vcenter_cluster
   datacenter_id = data.vsphere_datacenter.dc.id
 }

data "vsphere_resource_pool" "pool" {
  name          = "${var.vcenter_cluster}/Resources"
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_network" "network" {
  name = var.vcenter_network
  datacenter_id = data.vsphere_datacenter.dc.id
}

resource "vsphere_folder" "folder" {
  count = var.vcenter_folder == null ? 0 : 1
  path          = var.vcenter_folder
  type          = "vm"
  datacenter_id = data.vsphere_datacenter.dc.id
}
