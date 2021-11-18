data "aws_instance" "avi_controller" {
  count = var.controller_counts
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller-1"]
  }
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_current_password
  avi_controller =  data.aws_instance.avi_controller[0].public_ip
  avi_tenant     = "admin"
}

resource "avi_useraccount" "avi_user" {
  username     = var.avi_username
  old_password = var.avi_current_password
  password     = var.avi_new_password
}

