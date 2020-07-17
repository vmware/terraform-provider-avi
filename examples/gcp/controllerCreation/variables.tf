variable "credential_file" {
  type        = string
  description = "Enter the path to your credential file, for instance if my credential file is at \"~/Desktop/my_credential_file.json\", I would enter: ~/Desktop/my_credential_file.json"
}

variable "project_id" {
  type        = string
  description = "Enter your Project ID at which you want to create VM"
}

variable "region" {
  type        = string
  description = "Enter your region to manage resources in"
}

variable "bucket_name" {
  type        = string
  description = "(Optional) Enter the name of the bucket to which you want to upload *.tar.gz (image) file. If this variable's value isn't given, we will generate a bucket with random name"
  default     = "not-given"
}

variable "bucket_project" {
  type        = string
  description = "(Optional) Enter the project name to which you want the bucket to be created in case you haven't given bucket name different than default value. If you don't give value to this variable different than default, then your previous given project-id will be used."
  default     = "not-given"
}

variable "gcp_controller_dir" {
  type        = string
  description = "Enter the path to download gcp_controller.tar.gz file. For instance if the file is at \"./gcp_controller.tar.gz\", enter: ./gcp_controller.tar.gz"
}

variable "machine_type" {
  type        = string
  description = "(Optional) Machine type for your VM, should have atleast 8 cores and 24 gigs of memory"
  default     = "custom-8-24576"
  // default should have 8 core and 24 gigs of memory
}

variable "zone" {
  type        = string
  description = "Zone for your VM"
}

variable "network" {
  type        = string
  description = "Network for your VM"
}

variable "subnetwork" {
  type        = string
  description = "Subnetwork for your VM"
}

variable "subnetwork_project" {
  type        = string
  description = "(Optional) The project in which the subnetwork belongs. If this field is not provided, the given provider project id is used."
  default     = "not-given"
}

variable "email" {
  type        = string
  description = "(Optional) (For VM instance) The service account e-mail address. If not given, the default Google Compute Engine service account is used."
  default     = "not-given"
}

variable "boot_disk_size" {
  type        = string
  description = "The size of the image in gigabytes. If not specified, it will default to 50GB"
  default     = "50"
}

variable "controller_count" {
  type        = string
  description = "Number of controllers to create, enter only 1 or 3"
  default     = "1"
}

variable "network_ip" {
  type        = list
  description = "(Optional) The private IP address to assign to each instance. If empty, the address will be automatically assigned. Length should match with number of controllers"
  default     = []
}

variable "controller_name" {
  type        = string
  description = "Controller name"
  default     = "avi-controller"
}


variable "firewall_rules_name" {
  type        = string
  description = "We will be creating firewall rules as listed here: https://avinetworks.com/docs/18.2/avi-deployment-guide-for-google-cloud-platform-gcp/, following variable is just for its name."
}
