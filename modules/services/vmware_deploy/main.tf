
# ---------------------------------------------------------------------------------------------------------------------
# CONFIGURE THE VM FROM GIVING TEMPLATE
# ---------------------------------------------------------------------------------------------------------------------

data "vsphere_datacenter" "dc" {
  name = var.vm_datacenter
}

data "vsphere_datastore" "datastore" {
  name          = var.vm_datastore
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_resource_pool" "resource_pool" {
  name          = var.vm_resource_pool
  datacenter_id = data.vsphere_datacenter.dc.id
}


data "vsphere_network" "network" {
  name          = var.vm_network
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_virtual_machine" "template" {
  name          = var.vm_template
  datacenter_id = data.vsphere_datacenter.dc.id
}

resource "vsphere_virtual_machine" "vm" {
  name             = "${var.vm_name}-${count.index+1}"
  resource_pool_id = data.vsphere_resource_pool.resource_pool.id
  datastore_id     = data.vsphere_datastore.datastore.id
  count = 3
  num_cpus = 8
  memory = 24576
  guest_id = data.vsphere_virtual_machine.template.guest_id
  scsi_type = data.vsphere_virtual_machine.template.scsi_type

  folder = var.vm_folder
  network_interface {
    network_id   = data.vsphere_network.network.id
    adapter_type = data.vsphere_virtual_machine.template.network_interface_types[0]
  }

  disk {
    label = "disk1"
    size = data.vsphere_virtual_machine.template.disks.0.size
    unit_number = 0
    eagerly_scrub    = data.vsphere_virtual_machine.template.disks.0.eagerly_scrub
    thin_provisioned = data.vsphere_virtual_machine.template.disks.0.thin_provisioned
  }

  clone {
    template_uuid = data.vsphere_virtual_machine.template.id
  }
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = vsphere_virtual_machine.vm[0].default_ip_address
  avi_tenant     = "admin"
}

resource "avi_useraccount" "avi_user" {
  username     = var.avi_username
  old_password = var.avi_password
  password     = var.avi_new_password
}

resource "avi_cluster" "vmware_cluster" {
  name = "cluster-0-1"
  nodes {
    ip {
      type = "V4"
      addr = vsphere_virtual_machine.vm[0].default_ip_address
    }
    name = vsphere_virtual_machine.vm[0].name
  }
  nodes {
    ip {
      type = "V4"
      addr = vsphere_virtual_machine.vm[1].default_ip_address
    }
    name = vsphere_virtual_machine.vm[1].name
  }
  nodes {
    ip {
      type = "V4"
      addr = vsphere_virtual_machine.vm[2].default_ip_address
    }
    name = vsphere_virtual_machine.vm[2].name
  }
}

