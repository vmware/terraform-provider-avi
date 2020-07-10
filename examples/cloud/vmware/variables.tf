variable "avi_username" {                                                                                                                                               
  type    = string
  default = ""
}
variable "avi_password" {
  type    = string
  default = ""
}
variable "avi_controller" {
  type    = string
  default = ""
}
variable "avi_version" {
  type    = string
  default = "18.2.10"
}
variable "tenant" {
  type    = string
  default = "admin"
}
variable "cloud_name" {
  type    = string
  default = "tf_vmware_cloud"
}
variable "vcenter_username" {
  type    = string
  default = ""
}
variable "vcenter_password" {
  type    = string
  default = ""
}
variable "vcenter_datacenter" {
  type    = string
  default = ""
}
variable "vcenter_management_network" {
  type    = string
  default = ""
}
variable "vcenter_privilege" {
  type    = string
  default = "WRITE_ACCESS"
}
variable "vcenter_vcenter_url" {
  type    = string
  default = ""
}
variable "vcenter_license_tier" {
  type    = string
  default = "ENTERPRISE_18"
}
variable "vcenter_license_type" {
  type    = string
  default = "LIC_CORES"
}
