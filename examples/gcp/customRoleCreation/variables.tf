variable "credential_file" {
  type        = string
  description = "Please enter the path to your google service account's credential file, for instance if my credential file is at \"~/Desktop/my_credential_file.json\", I would enter: ~/Desktop/my_credential_file.json"
}

variable "network_role_id" {
  type        = string
  description = "ID of the custom role to manage network resources"
  default     = "avi.network"
}

variable "se_role_id" {
  type        = string
  description = "ID of the custom role to manage network resources"
  default     = "avi.se"
}

variable "storage_role_id" {
  type        = string
  description = "ID of the custom role to manage network resources"
  default     = "avi.storage"
}

variable "se_project_id" {
  type        = string
  description = "The project where the service account will be added for managing service engine resource"
}

variable "network_project_id" {
  type        = string
  description = "The project where the service account will be added for managing network resources"
}

variable "storage_project_id" {
  type        = string
  description = "The project where the service account will be added for managing storage resources"
}

variable "sa_account_project" {
  type        = string
  description = "Project name where service account to be assigned our custom roles is present"
}

variable "sa_account_id" {
  type        = string
  description = "The Service account id of the service account chosen to be assigned our custom roles"
}

