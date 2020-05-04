provider "vsphere" {
    user           = var.vsphare_user
    password       = var.vsphare_password
    vsphere_server = var.vsphere_server
    allow_unverified_ssl = true
//    version = "1.17.0"
}

data "vsphere_datacenter" "dc" {
  name = "${var.vsphere_datacenter}"
}

data "vsphere_datastore" "datastore" {
  name          = "${var.vm_datastore}"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}

data "vsphere_resource_pool" "pool" {
  name          = "${var.resource_pool}"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}

data "vsphere_network" "network" {
  name          = "${var.vm_network}"
  datacenter_id = "${data.vsphere_datacenter.dc.id}"
}

data "vsphere_content_library" "library" {
  name = "${var.content_library}"
}

data "vsphere_content_library_item" "item" {
  name       = "${var.vm_template}"
  library_id = "${data.vsphere_content_library.library.id}"
}

resource "vsphere_virtual_machine" "vm" {
  name             = "${var.vm_name}-${count.index+1}"
  resource_pool_id = "${data.vsphere_resource_pool.pool.id}"
  datastore_id     = "${data.vsphere_datastore.datastore.id}"
  count = 1
  num_cpus = 8
  memory = 24576

  folder = "${var.vm_folder}"
  network_interface {
    network_id   = "${data.vsphere_network.network.id}"
  }

  lifecycle {
    ignore_changes = ["guest_id"]
  }

  disk {
    label = "disk1"
    size = 130
    thin_provisioned = false
  }

  clone {
    template_uuid = "${data.vsphere_content_library_item.item.id}"
  }
}

