variable "avi_username" {
  type    = string
  default = "admin"
}

variable "avi_controller" {
  type    = string
  default = "10.206.114.103"
}
variable "avi_password" {
  type    = string
  default = "password"
}

variable "avi_tenant" {
  type    = string
  default = "admin"
}

variable "avi_cloud" {
  type    = string
  default = "NSX-T-TF-Cloud"
}

variable "avi_version" {
  type    = string
  default = "20.1.1"
}

variable "nsxt_url" {
  type    = string
  default = "10.206.113.90"
}

variable "nsxt_username" {
  type    = string
  default = "admin"
}

variable "nsxt_password" {
  type    = string
  default = "password"
}

variable "vsphere_server" {
  type    = string
  default = "10.206.113.91"
}

variable "vsphere_user" {
  type    = string
  default = "administrator@vsphere.local"
}

variable "vsphere_password" {
  type    = string
  default = "password"
}

variable "nsxt_cloudname" {
  type    = string
  default = "nsxt-tf-cloud"
}

variable "nsxt_cloud_prefix" {
  type    = string
  default = "swap-tf"
}

variable "transport_zone_name" {
  type    = string
  default = "nsx-overlay-transportzone"
}

variable "mgmt_lr_id" {
  type    = string
  default = "T1-Avi-MGMT"
}

variable "mgmt_segment_id" {
  type    = string
  default = "seg-avi-mgmt"
}

variable "data_lr_id" {
  type    = string
  default = "T1-A"
}

variable "data_segment_id" {
  type    = string
  default = "seg01A"
}

variable "vcenter_id" {
  type    = string
  default = "vc-01"
}

variable "content_library_name" {
  type    = string
  default = "Avi-SE-CL"
}
