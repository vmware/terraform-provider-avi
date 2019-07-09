# Configure the OpenStack Provider
provider "openstack" {
  user_name   = var.openstack_username
  tenant_name = var.openstack_tenant_name
  password    = var.openstack_password
  auth_url    = var.openstack_url
}

resource "openstack_compute_flavor_v2" "test-flavor" {
  name  = "avi.flavor"
  ram   = "8096"
  vcpus = "2"
  disk  = "80"
}

data "openstack_networking_network_v2" "network" {
  name = var.network_name
}

data "openstack_networking_subnet_v2" "terraform-subnets-0" {
  name = var.openstack_subnet
}

resource "openstack_images_image_v2" "avi-controller" {
  name             = "Avi-controller-18.1.3"
  local_file_path  = var.image_path
  container_format = "bare"
  disk_format      = "qcow2"
}

resource "openstack_compute_instance_v2" "avi-Openstack" {
  name              = "avi-op-instance"
  image_id          = openstack_images_image_v2.avi-controller.id
  flavor_id         = openstack_compute_flavor_v2.test-flavor.id
  availability_zone = "nova"

  network {
    uuid = data.openstack_networking_network_v2.network.id
  }
}

resource "openstack_compute_instance_v2" "avi-server-1" {
  name              = "pool-server-1"
  image_name        = "Jenkins-Server"
  flavor_id         = openstack_compute_flavor_v2.test-flavor.id
  availability_zone = "nova"

  network {
    uuid = data.openstack_networking_network_v2.network.id
  }
}

resource "openstack_compute_instance_v2" "avi-server-2" {
  name              = "pool-server-2"
  image_name        = "Jenkins-Server"
  flavor_id         = openstack_compute_flavor_v2.test-flavor.id
  availability_zone = "nova"

  network {
    uuid = data.openstack_networking_network_v2.network.id
  }
}

output "Instance Ip" {
  value = openstack_compute_instance_v2.avi-Openstack.network[0].fixed_ip_v4
}

output "Server-1 Ip" {
  value = openstack_compute_instance_v2.avi-server-1.network[0].fixed_ip_v4
}

output "Server-2 Ip" {
  value = openstack_compute_instance_v2.avi-server-2.network[0].fixed_ip_v4
}

output "Subnet Ip" {
  value = data.openstack_networking_subnet_v2.terraform-subnets-0.cidr
}

