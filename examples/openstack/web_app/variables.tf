variable "avi_username" {
  type    = "string"
  default = "admin"
}

variable "openstack_tenant_name" {
  type    = "string"
    default = "admin"
}

variable "openstack_username" {}

variable "openstack_password" {}

variable "openstack_url" {}

variable "avi_controller" {}

variable "avi_password" {}

variable "avi_tenant" {}

variable "avi_pool_server_1" {}

variable "avi_pool_server_2" {}

variable "network_name" {}

variable "op_subnet_ip" {}

variable "openstack_subnet" {}

variable "op_subnet_mask" {
  default = 21
}

