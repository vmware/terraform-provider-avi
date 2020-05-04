//Avi Provider variables

variable "vsphere_datacenter" {
  type    = string
  default = ""
}

variable "resource_pool" {
  type    = string
  default = ""
}

variable "content_library" {
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

variable "vsphare_user" {
  type = string
  default = ""
}

variable "vsphare_password" {
  type = string
  default = ""
}

variable "vsphere_server" {
  type = string
  default = ""
}

