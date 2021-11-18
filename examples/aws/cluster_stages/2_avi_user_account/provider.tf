terraform {
  required_providers {
    avi = {
      source = "vmware/avi"
      version = "21.1.1"
    }
  }
}
provider "aws" {
  shared_credentials_file = var.aws_creds_file
  region                  = var.aws_region
}


