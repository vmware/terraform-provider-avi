provider "vsphere" {
  user           = var.vsphere_user
  password       = var.vsphere_password
  vsphere_server = var.vsphere_server
  allow_unverified_ssl = true
  version = "1.17.4"
}

data "vsphere_datacenter" "dc" {
  name = var.vm_datacenter
}

data "vsphere_datastore" "datastore" {
  name          = var.vm_datastore
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_resource_pool" "pool" {
  name          = var.vm_resource_pool
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_network" "network" {
  name          = var.vm_network
  datacenter_id = data.vsphere_datacenter.dc.id
}

data "vsphere_content_library" "library" {
  name = var.content_library
}

data "vsphere_content_library_item" "item" {
  name       = var.vm_template
  library_id = data.vsphere_content_library.library.id
}

resource "vsphere_virtual_machine" "vm" {
  name             = "${var.vm_name}-${count.index+1}"
  resource_pool_id = data.vsphere_resource_pool.pool.id
  datastore_id     = data.vsphere_datastore.datastore.id
  count = 3
  num_cpus = 8
  memory = 24576
  folder = var.vm_folder
  network_interface {
    network_id   = data.vsphere_network.network.id
  }
  lifecycle {
    ignore_changes = [guest_id]
  }
  disk {
    label = "disk1"
    size = 130
    thin_provisioned = false
  }
  clone {
    template_uuid = data.vsphere_content_library_item.item.id
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

output "avi_cluster_output" {
  value = avi_cluster.vmware_cluster
}
