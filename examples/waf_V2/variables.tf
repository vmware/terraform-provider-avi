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
  default = ""
}
variable "pool_name" {
  type    = string
  default = "pool1"
}
variable "lb_algorithm" {
  type    = string
  default = "LB_ALGORITHM_ROUND_ROBIN"
}
variable "server1_ip" {
  type    = string
  default = "10.34.56.78"
}
variable "server1_port" {
  type    = number
  default = 8000
}
variable "poolgroup_name" {
  type    = string
  default = "poolgroup1"
}
variable "avi_terraform_vs_vip" {
  type    = string
  default = "10.20.23.45"
}
variable "vs_name" {
  type    = string
  default = "vs1"
}
variable "vs_port" {
  type    = number
  default = "8990"
}
variable "waf_crs_name" {
  type    = string
  default = "CRS-2021-1"
}
