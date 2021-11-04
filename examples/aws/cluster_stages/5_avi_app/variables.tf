variable "aws_access_key" {
  default = ""
}
variable "aws_secret_key" {
  default = ""
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
variable "aws_availability_zones" {
  default = ["us-west-2a", "us-west-2b", "us-west-2c"]
}
variable "aws_subnet_ip" {
  description = "An IP address from the subnet, example using 10.155.251.100 which is from the first network block"
}
variable "aws_subnet_mask" {
  default = 24
}
variable "project_name" {
}
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
variable "api_version" {
  default = "20.1.6"
}
variable "floating_ip" {
  default     = true
  description = "If you want to assign the floating IP to your VS"
}
