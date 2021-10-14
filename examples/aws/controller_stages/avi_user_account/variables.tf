variable "aws_region" {
  default = "us-west-2"
}

variable "avi_username" {
  default = "admin"
}

variable "avi_current_password" {
}

variable "avi_new_password" {
  default = "RemoMattei@21"
}

variable "project_name" {
}

variable "aws_creds_file" {
  default = "~/.aws/credentials"
}

variable "controller_counts" {
  default =  3

}
