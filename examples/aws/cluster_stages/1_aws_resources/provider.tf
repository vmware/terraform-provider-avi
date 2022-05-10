provider "aws" {
  #access_key = "${var.aws_access_key}"
  #secret_key = "${var.aws_secret_key}"
  shared_credentials_files = [var.aws_creds_file]
  region                   = var.aws_region
}
