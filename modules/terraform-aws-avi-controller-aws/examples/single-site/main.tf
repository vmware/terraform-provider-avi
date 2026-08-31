terraform {
  required_version = ">= 0.13.6"
  backend "local" {
  }
}
module "avi_controller_aws" {
  source = "../.."

  region                    = "us-west-2"
  aws_access_key            = var.aws_access_key
  aws_secret_key            = var.aws_secret_key
  create_networking         = var.create_networking
  create_iam                = var.create_iam
  avi_version               = var.avi_version
  custom_vpc_id             = var.custom_vpc_id
  custom_subnet_ids         = var.custom_subnet_ids
  avi_cidr_block            = var.avi_cidr_block
  controller_password       = var.controller_password
  key_pair_name             = var.key_pair_name
  private_key_path          = var.private_key_path
  name_prefix               = var.name_prefix
  controller_ha             = var.controller_ha
  controller_public_address = var.controller_public_address
  configure_dns_profile     = var.configure_dns_profile
  dns_service_domain        = var.dns_service_domain
  configure_dns_vs          = var.configure_dns_vs
  dns_vs_settings           = var.dns_vs_settings
}
