
provider "aws" {
  access_key = "${var.aws_access_key}"
  secret_key = "${var.aws_secret_key}"
  region     = "${var.aws_region}"
}

data "aws_instance" "avi_controller" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller"]
  }
}

data "aws_subnet" "terraform-subnets-0" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-subnet-0"]
  }
}

data "aws_subnet" "terraform-subnets-1" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-subnet-1"]
  }
}

data "aws_subnet" "terraform-subnets-2" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-subnet-2"]
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
  name         = "AWS"
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
      mgmt_network_name = "${data.aws_subnet.terraform-subnets-0.tags.Name}"
      mgmt_network_uuid = "${data.aws_subnet.terraform-subnets-0.id}"
      availability_zone = "${data.aws_subnet.terraform-subnets-0.availability_zone}"
    }

    zones = {
      mgmt_network_name = "${data.aws_subnet.terraform-subnets-1.tags.Name}"
      mgmt_network_uuid = "${data.aws_subnet.terraform-subnets-1.id}"
      availability_zone = "${data.aws_subnet.terraform-subnets-1.availability_zone}"
    }

    zones = {
      mgmt_network_name = "${data.aws_subnet.terraform-subnets-2.tags.Name}"
      mgmt_network_uuid = "${data.aws_subnet.terraform-subnets-2.id}"
      availability_zone = "${data.aws_subnet.terraform-subnets-2.availability_zone}"
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
