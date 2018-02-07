/**

All the AWS resources are created outside this plan in the aws_resources
directory.

Steps to add a new server
1. Add server in the ../aws_resources/avi_controller.tf

**/

provider "aws" {
  /*
                                                              Export the AWS credentials from the Environment. In order to explicitly
                                                              provide it in the plan then use the variables.tf to set aws_access_key and
                                                              aws_secret_key
                                                                $ export AWS_ACCESS_KEY_ID="anaccesskey"
                                                                $ export AWS_SECRET_ACCESS_KEY="asecretkey"
                                                                $ export AWS_DEFAULT_REGION="us-west-2"
                                                                */
  access_key = "${var.aws_access_key}"

  secret_key = "${var.aws_secret_key}"
  region     = "${var.aws_region}"
}

data "aws_instance" "avi_controller" {
  filter {
    name   = "tag:Name"
    values = ["grastogi-terraform-controller"]
  }
}

data "aws_subnet" "grastogi-terraform-subnet" {
  filter {
    name   = "tag:Name"
    values = ["grastogi-terraform-subnet"]
  }
}

data "aws_vpc" "avi_vpc" {
  id = "${var.aws_vpc_id}"
}

provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_password}"
  avi_controller = "${data.aws_instance.avi_controller.private_ip}"
  avi_tenant     = "admin"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

resource "avi_cloud" "aws_cloud_cfg" {
  name         = "aws_cloud_cfg"
  vtype        = "CLOUD_AWS"
  dhcp_enabled = true

  aws_configuration = {
    region            = "${var.aws_region}"
    secret_access_key = "${var.aws_secret_key}"
    access_key_id     = "${var.aws_access_key}"

    s3_encryption = {
      mode = "AWS_ENCRYPTION_MODE_NONE"
    }

    zones = {
      mgmt_network_name = "${data.aws_subnet.grastogi-terraform-subnet.tags.Name}"
      mgmt_network_uuid = "${data.aws_subnet.grastogi-terraform-subnet.id}"
      availability_zone = "${var.aws_availability_zone}"
    }

    use_iam_roles = false
    use_sns_sqs   = false
    vpc           = "${data.aws_vpc.avi_vpc.tags.Name}"
    vpc_id        = "${var.aws_vpc_id}"
  }

  license_tier = "ENTERPRISE_18"
  license_type = "LIC_CORES"
  tenant_ref   = "${data.avi_tenant.default_tenant.id}"
}
