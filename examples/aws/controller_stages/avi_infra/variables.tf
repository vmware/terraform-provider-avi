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
  default = "remo-demo"
}
