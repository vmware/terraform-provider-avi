variable "avi_password" {
  type    = "string"
  default = "admin"
}

variable "avi_controller" {
  type = "string"
  default = ""
}

variable "avi_version" {
  type = "string"
  default = "18.2.5"
}

variable "avi_app_server_v1" {
  type = "string"
  default = "10.90.64.63"
}

variable "avi_app_server_v2" {
  type = "string"
  default = "10.90.64.63"
}