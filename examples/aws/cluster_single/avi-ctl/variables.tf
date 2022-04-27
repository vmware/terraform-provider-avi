variable "old_password" {}

variable "iam_profile" {
  default = "remo-avi-controller"
}
variable "vpc_id_fse" {
  description = "VPC ID"
}
variable "admin_username" {
  description = "User Admin"
}
variable "admin_password" {
  description = "Admin Password"
}
variable "se_subnet_id" {
  description = "Subnet ID"
}
# variable "ami-image" {
#   type        = "string"
#   description = "Image to use"
# }
variable "image-size" {
  description = "Image size"
}
variable "controller_name" {
  description = "Controller Name"
}
variable "public_ip" {
  description = "If you want a Public IP"
}
variable "shutdown_rules" {
  description = "Add noshut if you want to keep it up"
}
variable "department_name" {
  description = "Department Name"
}
# this defines which AMI will be used
variable "myregion" {
  description = "Region"
}
variable "avi_sg_name" {
  description = "Name of the Sec Groups"
}
variable "profile" {
  description = "Profile to use "
}
variable "owner" {
  description = "owner"
}
# You could spin up more than 1 but use the other repo for cluster
variable "controller_count" {
  default = 1
}
variable "aws_access_key" {}
variable "aws_secret_key" {}
variable "api_version" {
  default = "21.1.4"
}
variable "dns_ip" {
  type    = list(string)
  default = ["8.8.8.8","8.8.4.4"]
}
variable "ntp_servers" {
  type    = list(string)
  default = ["0.us.pool.ntp.org","1.us.pool.ntp.org","2.us.pool.ntp.org","3.us.pool.ntp.org"]
}

variable "search_default_domain" {
  default = "io.local"
}
variable "trusted_cidr" {
  description = "cidr block of addresses trusted for access to the kubernetes API server"
}
variable "welcome_banner" {
  default = "Ciao and Welcome to my new Avi Controller!"
}
variable "tag_name" {
  default = "avi-ctl-tf"
}
