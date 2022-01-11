variable "aws_access_key" {
}

variable "aws_secret_key" {
}

variable "aws_region" {
  type    = string
  default = "us-west-2"
}

variable "aws_vpc_id" {
  type    = string
  default = "vpc-46246520"
}

variable "avi_username" {
  type    = string
  default = "admin"
}

variable "avi_password" {
}

variable "aws_availability_zones" {
  type    = list(string)
  default = ["us-west-2a", "us-west-2b", "us-west-2c"]
}

variable "project_name" {
  default = "awesome-avi"
}
variable "api_version" {
  default = "21.1.2"
}
variable "connection_mem_percentage" {
  default     = 50
  description = "default"
}

variable "disk_per_se" {
  default     = 10
  description = "default"
}
variable "ha_mode" {
  default = "HA_MODE_SHARED"
}

variable "instance_flavor_se" {
  default = "t2.large"
}
variable "max_se" {
  default     = 2
  description = ""
}
variable "max_vs_per_se" {
  default     = 20
  description = "max se for this group"
}
variable "mem_per_se" {
  default     = 2048
  description = "for WAF you will need more ram"
}
variable "vcpus_per_se" {
  default     = 2
  description = "How many CPU per SE"
}
variable "deprovision_delay" {
  default = 5
}
variable "se_prefix" {
  default = "Avi_SE_AA"
}
variable "enable_route53" {
  default = true
}
