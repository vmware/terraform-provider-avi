locals {
  # AKO Settings
  cloud_settings = {
    se_mgmt_subnets                 = var.create_networking ? local.mgmt_subnets : local.custom_mgmt_subnets
    vpc_id                          = var.create_networking ? aws_vpc.avi[0].id : var.custom_vpc_id
    aws_region                      = var.region
    avi_version                     = var.avi_version
    dns_servers                     = var.dns_servers
    dns_search_domain               = var.dns_search_domain
    ntp_servers                     = var.ntp_servers
    email_config                    = var.email_config
    se_name_prefix                  = var.name_prefix
    mgmt_security_group             = aws_security_group.avi_se_mgmt_sg.id
    data_security_group             = aws_security_group.avi_data_sg.id
    controller_ha                   = var.controller_ha
    controller_ip                   = local.controller_ip
    controller_names                = local.controller_names
    configure_dns_route_53          = var.configure_dns_route_53
    configure_dns_profile           = var.configure_dns_profile
    dns_service_domain              = var.dns_service_domain
    configure_dns_vs                = var.configure_dns_vs
    dns_vs_settings                 = var.dns_vs_settings
    configure_gslb                  = var.configure_gslb
    configure_gslb_additional_sites = var.configure_gslb_additional_sites
    gslb_site_name                  = var.gslb_site_name
    gslb_domains                    = var.gslb_domains
    additional_gslb_sites           = var.additional_gslb_sites
    se_ha_mode                      = var.se_ha_mode
  }
  controller_names = aws_instance.avi_controller[*].tags.Name
  controller_ip    = aws_instance.avi_controller[*].private_ip

  mgmt_subnets = { for subnet in aws_subnet.avi : subnet.availability_zone =>
    {
      "mgmt_network_uuid" = subnet.id
      "mgmt_network_name" = subnet.tags["Name"]
    }
  }
  custom_mgmt_subnets = { for subnet in data.aws_subnet.custom : subnet.availability_zone =>
    {
      "mgmt_network_uuid" = subnet.id
      "mgmt_network_name" = subnet.tags["Name"]
    }
  }
  az_names = data.aws_availability_zones.azs.names
  avi_ami = {
    "21.1.1" = {
      "us-east-1" = "ami-010074db16c05a098"
      "us-east-2" = "ami-063514c0bcbe293f9"
      "us-west-1" = "ami-03baad459ee4a3980"
      "us-west-2" = "ami-007edf6b2b302cce5"
    }
    "20.1.6" = {
      "us-east-1" = "ami-0645a1b3f7a3b0935"
      "us-east-2" = "ami-09605d05d713e13f0"
      "us-west-1" = "ami-0bd87553e8167e589"
      "us-west-2" = "ami-0451e26f70764fc9e"
    }
    "20.1.5" = {
      "us-east-1" = "ami-0c005433aaab63ab8"
      "us-east-2" = "ami-0426b8de85eb12a6a"
      "us-west-1" = "ami-0d4e88b10724f412f"
      "us-west-2" = "ami-0c6aceb66c9e18cd6"
    }
    "20.1.4" = {
      "us-east-1" = "ami-0822ac66b62a893cc"
      "us-east-2" = "ami-08cf08a0ba5dcff76"
      "us-west-1" = "ami-0fa6c926c2fb340d3"
      "us-west-2" = "ami-08ecc90c68f430fc2"
    }
    "20.1.3" = {
      "us-east-1" = "ami-0ca5d1b0c6e2ef1f8"
      "us-east-2" = "ami-0786e4ed20b179355"
      "us-west-1" = "ami-064facee296f10111"
      "us-west-2" = "ami-0f1ba8c00dfd6c04c"
    }
    "20.1.2" = {
      "us-east-1" = "ami-090699c2c811f2960"
      "us-east-2" = "ami-0a2adfbfbd4c6da8e"
      "us-west-1" = "ami-0d440bb0a05717a19"
      "us-west-2" = "ami-08a85624884915308"
    }
    "18.2.11" = {
      "us-east-1" = "ami-071db10cf4818c77a"
      "us-east-2" = "ami-0858e11a36f912dcd"
      "us-west-1" = "ami-071db10cf4818c77a"
      "us-west-2" = "ami-07963584f136c2e33"
    }
  }
}
resource "aws_instance" "avi_controller" {
  count = var.controller_ha ? 3 : 1
  ami   = local.avi_ami[var.avi_version][var.region]
  root_block_device {
    volume_size           = var.boot_disk_size
    delete_on_termination = true
  }
  instance_type = var.instance_type
  key_name      = var.key_pair_name
  #availability_zone = var.create_networking ? aws_subnet.avi[count.index].availability_zone : 
  subnet_id              = var.create_networking ? aws_subnet.avi[count.index].id : var.custom_subnet_ids[count.index]
  vpc_security_group_ids = [aws_security_group.avi_controller_sg.id]
  iam_instance_profile   = var.create_iam ? aws_iam_instance_profile.avi[0].id : null
  tags = {
    Name = "${var.name_prefix}-avi-controller-${count.index + 1}"
  }
  lifecycle {
    ignore_changes = [tags]
  }
}
resource "aws_ec2_tag" "custom_controller_1" {
  for_each    = var.custom_tags
  resource_id = aws_instance.avi_controller[0].id
  key         = each.key
  value       = each.value
}
resource "aws_ec2_tag" "custom_controller_2" {
  for_each    = var.controller_ha ? var.custom_tags : {}
  resource_id = aws_instance.avi_controller[1].id
  key         = each.key
  value       = each.value
}
resource "aws_ec2_tag" "custom_controller_3" {
  for_each    = var.controller_ha ? var.custom_tags : {}
  resource_id = aws_instance.avi_controller[2].id
  key         = each.key
  value       = each.value
}
resource "null_resource" "ansible_provisioner" {
  # Changes to any instance of the cluster requires re-provisioning
  triggers = {
    controller_instance_ids = join(",", aws_instance.avi_controller.*.id)
  }

  connection {
    type        = "ssh"
    host        = var.controller_public_address ? aws_instance.avi_controller[0].public_ip : aws_instance.avi_controller[0].private_ip
    user        = "admin"
    timeout     = "600s"
    private_key = file(var.private_key_path)
  }
  provisioner "file" {
    content = templatefile("${path.module}/files/avi-cleanup.yml.tpl",
    local.cloud_settings)
    destination = "/home/admin/avi-cleanup.yml"
  }
  provisioner "file" {
    content = templatefile("${path.module}/files/avi-controller-aws-all-in-one-play.yml.tpl",
    local.cloud_settings)
    destination = "/home/admin/avi-controller-aws-all-in-one-play.yml"
  }
  provisioner "remote-exec" {
    inline = [
      "sleep 180",
      "sudo /opt/avi/scripts/initialize_admin_user.py --password ${var.controller_password}",
      "ansible-playbook avi-controller-aws-all-in-one-play.yml -e password=${var.controller_password} -e aws_access_key_id=${var.aws_access_key} -e aws_secret_access_key=${var.aws_secret_key} > ansible-playbook.log 2> ansible-error.log",
      "echo Controller Configuration Completed"
    ]
  }
}
