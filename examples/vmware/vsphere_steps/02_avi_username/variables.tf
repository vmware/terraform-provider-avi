variable "avi_version" {}
variable "avi_controller_ips" {}
variable "avi_tenant" {}
variable "avi_current_password" {
  sensitive = true
}
variable "avi_password" {
  sensitive = true
}
variable "avi_tenant" {
  default = "admin"
}