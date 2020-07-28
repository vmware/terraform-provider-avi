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
variable "lb_algorithm_consistent_hash_hdr_var" {
  default = null
}
variable "pools" {
  description = "map of trap servers for sending alarms via snmp traps"
  type = any
  default = []
}
