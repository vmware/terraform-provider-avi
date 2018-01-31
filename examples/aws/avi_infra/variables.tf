variable "aws_access_key" {
  type    = "string"
  default = ""
}

variable "aws_secret_key" {
  type    = "string"
  default = ""
}

variable "aws_region" {
  type    = "string"
  default = "us-west-2"
}

variable "aws_vpc_id" {
  type    = "string"
  default = "vpc-19295f7c"
}

variable "avi_username" {
  type    = "string"
  default = "admin"
}

variable "avi_password" {}

variable "aws_availability_zone" {
  type    = "string"
  default = "us-west-2a"
}
