variable "avi_username" {
}
variable "avi_password" {
}
variable "avi_controller" {
  default = "100.20.43.30"
}
variable "avi_version" {
  default = "20.1.6"
}
variable "avi_tenant" {
  default = "admin"
}
variable "waf_name" {
  default = "System-WAF-Policy"
}
variable "cloud_name" {
  default = "AWS"
}
variable "se_prefix" {
  default = "WAF_SE"
}
variable "accelerated_networking" {
  default = true
}
variable "inst_flavor" {
  default = "c4.xlarge"
}
variable "hyperthreaded_cores" {
  default = false
}
variable "se_group_name" {
  default = "Default-Group"
}
