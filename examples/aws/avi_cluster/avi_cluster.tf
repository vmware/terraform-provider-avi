provider "aws" {
  shared_credentials_file = "${var.aws_creds_file}"
  region     = "${var.aws_region}"
}

data "aws_instance" "avi_controller" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller"]
  }
}


data "aws_instance" "avi_controller_2" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller-2"]
  }
}

data "aws_instance" "avi_controller_3" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller-3"]
  }
}

resource "avi_cluster" "aws_cluster" {
  name = "cluster-0-1"
  nodes {
    ip = {
      type = "V4"
      addr = "${data.aws_instance.avi_controller.private_ip}" }
    name = "node01"
  }
  nodes{
    ip = {
      type = "V4"
      addr = "${data.aws_instance.avi_controller_2.private_ip}" }
    name = "node02"
  }
  nodes{
    ip = {
      type = "V4"
      addr = "${data.aws_instance.avi_controller_3.private_ip}" }
    name = "node03"
  }
}

provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_password}"
  avi_controller = "${data.aws_instance.avi_controller.private_ip}"
  avi_tenant     = "admin"
}
