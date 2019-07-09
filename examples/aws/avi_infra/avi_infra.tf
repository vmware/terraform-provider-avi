provider "aws" {
  access_key = var.aws_access_key
  secret_key = var.aws_secret_key
  region     = var.aws_region
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
  id = var.aws_vpc_id
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = data.aws_instance.avi_controller.private_ip
  avi_tenant     = "admin"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

resource "avi_cloud" "aws_cloud_cfg" {
  name         = "AWS"
  vtype        = "CLOUD_AWS"
  dhcp_enabled = true

  aws_configuration {
    region              = var.aws_region
    secret_access_key   = var.aws_secret_key
    access_key_id       = var.aws_access_key
    route53_integration = true
    s3_encryption {
      mode = "AWS_ENCRYPTION_MODE_NONE"
    }

    zones {
      mgmt_network_name = data.aws_subnet.terraform-subnets-0.tags.Name
      mgmt_network_uuid = data.aws_subnet.terraform-subnets-0.id
      availability_zone = data.aws_subnet.terraform-subnets-0.availability_zone
    }

    zones {
      mgmt_network_name = data.aws_subnet.terraform-subnets-1.tags.Name
      mgmt_network_uuid = data.aws_subnet.terraform-subnets-1.id
      availability_zone = data.aws_subnet.terraform-subnets-1.availability_zone
    }

    zones {
      mgmt_network_name = data.aws_subnet.terraform-subnets-2.tags.Name
      mgmt_network_uuid = data.aws_subnet.terraform-subnets-2.id
      availability_zone = data.aws_subnet.terraform-subnets-2.availability_zone
    }

    use_iam_roles = false
    use_sns_sqs   = false
    vpc           = data.aws_vpc.avi_vpc.tags.Name
    vpc_id        = var.aws_vpc_id
  }

  license_tier = "ENTERPRISE_18"
  license_type = "LIC_CORES"
  tenant_ref   = data.avi_tenant.default_tenant.id
}

resource "avi_serviceenginegroup" "aws_se_group" {
  name                         = "Default-Group"
  archive_shm_limit            = 8
  algo                         = "PLACEMENT_ALGO_PACKED"
  archive_shm_limit            = 8
  buffer_se                    = 0
  cloud_ref                    = avi_cloud.aws_cloud_cfg.id
  connection_memory_percentage = 50
  disk_per_se                  = 10
  enable_vip_on_all_interfaces = true
  ha_mode                      = "HA_MODE_SHARED"
  instance_flavor              = "t2.large"
  license_tier                 = "ENTERPRISE_18"
  license_type                 = "LIC_CORES"
  se_bandwidth_type            = "SE_BANDWIDTH_UNLIMITED"
  max_se                       = 2
  max_vs_per_se                = 20
  memory_per_se                = 2048
  min_scaleout_per_vs          = 1
  realtime_se_metrics {
    duration = 0
    enabled  = true
  }
  vcpus_per_se         = 2
  se_deprovision_delay = 5
  se_name_prefix       = var.project_name
  tenant_ref           = data.avi_tenant.default_tenant.id
}

