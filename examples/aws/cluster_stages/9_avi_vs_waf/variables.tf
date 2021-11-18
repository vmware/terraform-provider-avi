variable "avi_username" {
  type    = string
  default = "admin"
}
variable "avi_password" {
  type    = string
  default = "NationwidePOC"
}
variable "avi_controller" {
  type    = string
  default = "10.0.21.58"
}
variable "avi_version" {
  type    = string
  default = "20.1.5"
}
variable "vs_name" {
  type    = string
  default = "nwpoc-waf"
}
variable "vs_fqdn" {
  type    = string
  default = "nwpoc-waf.wafvpc.nwpoc.demoavi.us"
}
variable "vs_cloud" {
  type    = string
  default = "WAF-VPC"
}
