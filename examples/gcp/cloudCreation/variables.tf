variable "credential_file" {
  type        = string
  description = "Please enter the path to your google service account's credential file, for instance if my credential file is at \"~/Desktop/my_credential_file.json\", I would enter: ~/Desktop/my_credential_file.json"
}

variable "controller_count" {
  type        = string
  description = "Number of Controller instances (should be either 1 or 3)"
}

variable "project_id" {
  type        = string
  description = "The ID of the project in which the controller's belong."
}

variable "se_project_id" {
  type        = string
  description = "Google Cloud Platform Project ID where Service Engines will be spawned."
}

variable "region_name" {
  type        = string
  description = "Google Cloud Platform Region Name where Service Engines will be spawned."
}

variable "zone" {
  type        = string
  description = "Zone of your VM"
}

variable "avi_username" {
  type        = string
  description = "Username of your Avi account"
}

variable "avi_current_password" {
  type        = string
  description = "Current password of your Avi account"
}

variable "avi_version" {
  type        = string
  description = "Avi version"
}

variable "controller_name" {
  type        = string
  description = "Controller name"
  default     = "avi-controller"
}

variable "vpc_project_id" {
  type        = string
  description = "Enter VPC Project ID"
}

variable "vpc_network" {
  type        = string
  description = "Enter VPC Network"
}

variable "vpc_subnetwork" {
  type        = string
  description = "Enter VPC Subnetwork"
}

variable "zones" {
  type        = list(string)
  description = "Google Cloud Platform Zones where Service Engines will be distributed for HA."
}

variable "gcp_cloud_name" {
  type        = string
  description = "Enter name for GCP Cloud object"
}

variable "firewall_target_tags" {
  type        = list(string)
  description = "Firewall rule network target tags which will be applied on Service Engines to allow ingress and egress traffic for Service Engines."
  default     = ["http-server", "https-server"]
}

variable "gcp_user" {
  type        = string
  description = "Name of cloud connector user"
}

variable "gcs_bucket_name" {
  type        = string
  description = "(Optional) Google Cloud Storage Bucket Name where Service Engine image will be uploaded. This image will be deleted once the image is created in Google compute images. By default, a bucket will be created if this field is not specified."
  default     = "not-given"
}

variable "gcs_project_id" {
  type        = string
  description = "(Optional) Google Cloud Storage Project ID where Service Engine image will be uploaded. This image will be deleted once the image is created in Google compute images. By default, se_project_id will be used."
  default     = "not-given"
}

variable "tenant" {
  type        = string
  description = "Name of avi tenant"
  default = "admin"
}
