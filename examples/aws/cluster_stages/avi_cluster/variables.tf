variable "aws_region" {
  default = "us-west-2"
}

variable "aws_vpc_id" {
  default = "vpc-19295f7c"
}

variable "avi_username" {
  default = "admin"
}

variable "avi_password" {
}

variable "aws_availability_zones" {
  default = ["us-west-2a", "us-west-2b", "us-west-2c"]
}

variable "project_name" {
  default = "remo-demo"
}

variable "aws_creds_file" {
  default = "~/.aws/credentials"
}

variable "cluster_name" {
  default = "avi-cluster-tf"
}
variable "api_version" {
  default = "20.1.6"
}
