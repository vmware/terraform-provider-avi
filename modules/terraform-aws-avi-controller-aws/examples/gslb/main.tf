terraform {
  required_version = ">= 0.13.6"
  backend "local" {
  }
}

module "avi_controller_aws_west2" {
  source = "../.."

  region                    = "us-west-2"
  aws_access_key            = var.aws_access_key
  aws_secret_key            = var.aws_secret_key
  controller_ha             = var.controller_ha
  controller_public_address = var.controller_public_address
  create_networking         = var.create_networking
  create_iam                = var.create_iam
  avi_version               = var.avi_version
  custom_vpc_id             = var.custom_vpc_id_west
  custom_subnet_ids         = var.custom_subnet_ids_west
  avi_cidr_block            = var.avi_cidr_block_west
  controller_password       = var.controller_password
  key_pair_name             = var.key_pair_name
  private_key_path          = var.private_key_path
  name_prefix               = var.name_prefix_west
  configure_dns_profile     = "true"
  dns_service_domain        = "west1.avidemo.net"
  configure_dns_vs          = "true"
  dns_vs_settings           = var.dns_vs_settings_west
}
module "avi_controller_aws_east2" {
  source = "../.."

  region                          = "us-east-2"
  aws_access_key                  = var.aws_access_key
  aws_secret_key                  = var.aws_secret_key
  controller_ha                   = var.controller_ha
  controller_public_address       = var.controller_public_address
  create_networking               = var.create_networking
  create_iam                      = var.create_iam
  avi_version                     = var.avi_version
  custom_vpc_id                   = var.custom_vpc_id_east
  custom_subnet_ids               = var.custom_subnet_ids_east
  avi_cidr_block                  = var.avi_cidr_block_east
  controller_password             = var.controller_password
  key_pair_name                   = var.key_pair_name
  private_key_path                = var.private_key_path
  name_prefix                     = var.name_prefix_east
  configure_dns_profile           = "true"
  dns_service_domain              = "east2.avidemo.net"
  configure_dns_vs                = "true"
  dns_vs_settings                 = var.dns_vs_settings_east
  configure_gslb                  = "true"
  gslb_site_name                  = "East2"
  gslb_domains                    = ["gslb.avidemo.net"]
  configure_gslb_additional_sites = "true"
  additional_gslb_sites           = [{ name = "West2", ip_address_list = module.avi_controller_aws_west2.controllers[*].private_ip_address, dns_vs_name = "DNS-VS" }]
}
