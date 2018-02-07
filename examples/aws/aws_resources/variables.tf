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

variable "aws_subnet_ip" {
  default = "10.144.44.0"
}

variable "aws_subnet_mask" {
  default = 24
}

variable "aws_available_zone" {
  default = "us-west-2a"
}
