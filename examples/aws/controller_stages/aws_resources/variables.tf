variable "aws_access_key" {
  default = ""
}

variable "aws_secret_key" {
  default = ""
}

variable "controller_counts" {
  default = 3

}
variable "aws_region" {
  default = "us-west-2"
}

variable "aws_vpc_id" {
}

variable "avi_username" {
  default = "admin"
}

variable "avi_password" {
}

variable "aws_availability_zone" {
  default = "us-west-2a"
}

variable "aws_subnet_mask" {
  default = 24
}

variable "project_name" {
  default = "remo-demo"
}

variable "aws_creds_file" {
  default = "~/.aws/credentials"
}

variable "aws_subnets" {
  default = ["10.155.251.0", "10.155.252.0", "10.155.253.0"]
}

variable "avi_controller_iam" {
  default = "remo-avi-controller"
}
variable "avi_controller_ami" {
  default = "ami-0451e26f70764fc9e"
}

variable "aws_availability_zones" {
  type    = list(string)
  default = ["us-west-2a", "us-west-2b", "us-west-2c"]
}

