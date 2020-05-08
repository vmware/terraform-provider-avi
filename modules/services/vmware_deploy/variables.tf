variable "vm_datacenter" {
  type    = string
  default = ""
}

variable "vm_resource_pool" {
  type    = string
  default = ""
}

variable "vm_datastore" {
  type    = string
  default = ""
}

variable "vm_name" {
  type    = string
  default = ""
}

variable "vm_network" {
  type    = string
  default = ""
}

variable "vm_folder" {
  type    = string
  default = ""
}

variable "vm_template" {
  type    = string
  default = ""
}

// Avi Provider variables
variable "avi_username" {
  type = string
  default = ""
}

variable "avi_password" {
  type = string
  default = ""
}

variable "avi_new_password" {
  type = string
  default = ""
}
