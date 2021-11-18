provider "aws" {
  #access_key = var.aws_access_key
  #secret_key = var.aws_secret_key
  shared_credentials_file = var.aws_creds_file
  region                  = var.aws_region
}
provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = data.aws_instance.avi_controller.public_ip
  avi_tenant     = var.avi_tenant
  avi_version    = var.api_version
}


