// GCP Variables

variable "credential_file" {
  type        = string
  description = "Please enter the path to your credential file, for instance if my credential file is at \"~/Desktop/my_credential_file.json\", I would enter: ~/Desktop/my_credential_file.json"
}

variable "project_id" {
  type        = string
  description = "The ID of the project in which the VM's belongs"
}

variable "zone" {
  type        = string
  description = "Zone of your VM"
}

variable "controller_name" {
  type        = string
  description = "Controller name"
  default     = "avi-controller"
}

// Avi Variables

variable "avi_username" {
  type        = string
  description = "Username for your Avi account"
}

variable "avi_current_password" {
  type        = string
  description = "Current password of your Avi account"
}

variable "avi_version" {
  type        = string
  description = "Avi version"
}

variable "avi_new_password" {
  type        = string
  description = "New password for your Avi account"
}

variable "cluster_name" {
  type        = string
  description = "Cluster name"
  default     = "avi-cluster"
}

variable "tenant" {
  type        = string
  description = "Name of avi tenant"
  default = "admin"
}
