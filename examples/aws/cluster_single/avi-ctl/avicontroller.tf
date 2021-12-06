data "template_file" "userdata" {
  template = file("files/userdata.json")
  vars = {
    avi_controller_password = var.admin_password
    dns                     = var.dns_ip
    dns1                    = var.dns1_ip
    search_domain           = var.search_default_domain
    welcome_banner          = var.welcome_banner
  }
}
resource "aws_instance" "avi-controller" {
  count                       = var.controller_count
  user_data                   = count.index == 0 ? data.template_file.userdata.rendered : ""
  ami                         = lookup(var.ami-image, var.myregion)
  associate_public_ip_address = var.public_ip
  instance_type               = var.image-size
  subnet_id                   = var.se_subnet_id
  key_name                    = aws_key_pair.generated.key_name
  iam_instance_profile        = var.iam_profile
  vpc_security_group_ids = [
    aws_security_group.remo_sg.id,
  ]
  tags = {
    Name            = "ctl-avi-tf-${count.index + 1}"
    dept            = var.department_name
    shutdown_policy = var.shutdown_rules
    owner           = var.owner
  }
}

