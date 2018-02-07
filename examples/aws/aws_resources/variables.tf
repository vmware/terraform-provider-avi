variable "aws_access_key" {}

variable "aws_secret_key" {}

variable "aws_region" {
  type    = "string"
  default = "us-west-2"
}

variable "aws_vpc_id" {
  type    = "string"
  default = "vpc-19295f7c"
}

variable "aws_subnets" {
  default = ["10.144.43.0", "10.144.44.0", "10.144.45.0"]
}

variable "aws_subnet_mask" {
  default = 24
}

variable "aws_available_zone" {
  default = "us-west-2a"
}

variable "project_name" {}

variable "avi_controller_ami" {
  default = "ami-2c0bbf54"
}

variable "aws_availability_zones" {
  default = ["us-west-2a", "us-west-2b", "us-west-2c"]
}
