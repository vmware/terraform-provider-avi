data "template_file" "userdata" {
  template = file("files/userdata.json")
  vars = {
    password = var.admin_password
  }
}
resource "aws_instance" "avi-controller" {
  count                       = var.controller_count
  user_data                   = "${count.index ==0 ? data.template_file.userdata.rendered :"" }"
  ami                         = lookup(var.ami-image, var.myregion)
  associate_public_ip_address = var.public_ip
  instance_type               = var.image-size
  subnet_id                   = var.se_subnet_id
  key_name                    = aws_key_pair.generated.key_name
  iam_instance_profile        = var.iam_profile
  vpc_security_group_ids      = [
      aws_security_group.remo_sg.id,
  ]
  tags                        = {
    Name                      = "ctl-avi-tf-rm-${count.index +1}"
    dept                      = var.department_name
    shutdown_policy           = var.shutdown_rules
    owner                     = var.owner
  }
}

terraform {
  required_providers {
    avi = {
      source = "vmware/avi"
      version = "21.1.1"
    }
  }
}
provider "avi" {
  avi_username   = var.admin_username
  avi_password   = var.admin_password
  avi_controller = aws_instance.avi-controller[0].public_ip
  avi_tenant     = "admin"
  avi_version     = var.api_version
}
