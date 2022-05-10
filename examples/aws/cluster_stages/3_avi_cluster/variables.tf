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
}

variable "aws_creds_file" {
  default = "~/.aws/credentials"
}

variable "cluster_name" {
  default = "avi-cluster-tf"
}
variable "api_version" {
  default = "21.1.4"
}
variable "ntp_servers" {
  type    = list(any)
  default = ["0.us.pool.ntp.org", "1.us.pool.ntp.org", "2.us.pool.ntp.org", "3.us.pool.ntp.org"]
}
variable "mail_server_tls" {
  default = false
}
variable "email" {
  default = "admin@avinetworks.com"
}
variable "mail_server" {
  default = "localhost"
}
variable "mail_server_port" {
  default = 25
}
variable "mail_type" {
  default = "SMTP_LOCAL_HOST"
}
variable "dns_servers" {
  default = ["8.8.8.8", "8.8.4.4", "1.1.1.1"]
}
variable "search_domain" {
  default = "io.local"
}
variable "banner" {
  default = "Avi Demo with Terraform"
}
