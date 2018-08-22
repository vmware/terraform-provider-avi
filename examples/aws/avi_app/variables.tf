variable "aws_access_key" {
  default=""
}

variable "aws_secret_key" {
  default =""
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

variable "aws_availability_zones" {
  type = "list"
  default = ["us-west-2a", "us-west-2b", "us-west-2c"]
}

variable "aws_subnet_ip" {
  default = ""
}

variable "aws_subnet_mask" {
  default = 24
}

variable "project_name" {}

variable "webserver_ami" {
  default = "ami-05141f7c"
}

variable "webserver_instance_type" {
  default = "t2.micro"
}

variable "webserver_count" {
  default = 4
}

variable "aws_creds_file" {
  default = "~/.aws/credentials"
}