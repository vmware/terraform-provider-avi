variable "avi_username" {
}
variable "avi_password" {
}
variable "avi_controller" {
}
variable "avi_version" {
  default = "21.1.2"
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
  default = "SE_Group_One"
}
variable "csr_version" {
  default  = "CRS-2021-2" # 21.1.4
  #default  = "CRS-2021-1" # 21.1.3
  #default = "CRS-2020-3" #20.1.6
}
variable "owner_key" {
  default = "Owner"
}
variable "owner_value" {
  default = "Your Name"
}
variable "policy_key" {
  default = "shutdown_policy"
}
variable "policy_value" {
  default = "noshut"
}
variable "dept_key" {
  default = "Dept"
}
variable "dept_value" {
  default = "My TEAM"
}
