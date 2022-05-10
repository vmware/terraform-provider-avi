//
// Usage:
//  terraform apply -var project_name=<xx> -var avi_password=<xx> -var aws_access_key=<xx> -var aws_secret_key=<xx> -var aws_subnet_ip=<xx>
//


data "avi_applicationprofile" "system_http_profile" {
  name = "System-HTTP"
}

data "avi_applicationprofile" "system_https_profile" {
  name = "System-Secure-HTTP"
}
data "aws_instance" "avi_controller" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-controller-1"]
  }
}
resource "aws_instance" "terraform-webserver" {
  count         = var.webserver_count
  ami           = var.webserver_ami
  instance_type = var.webserver_instance_type
  subnet_id     = data.aws_subnet.terraform-subnets-0.id
  tags = {
    Name    = "${var.project_name}-terraform-webserver-${count.index}"
    Project = "${var.project_name}-terraform-webservers"
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

data "avi_tenant" "default_tenant" {
  name = "admin"
}

data "avi_cloud" "aws_cloud_cfg" {
  name = "AWS"
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = data.avi_cloud.aws_cloud_cfg.id
}

data "avi_healthmonitor" "system_http_healthmonitor" {
  name = "System-HTTP"
}

data "avi_networkprofile" "system_tcp_profile" {
  name = "System-TCP-Proxy"
}

data "avi_analyticsprofile" "system_analytics_profile" {
  name = "System-Analytics-Profile"
}

data "avi_sslkeyandcertificate" "system_default_cert" {
  name = "System-Default-Cert"
}

data "avi_sslprofile" "system_standard_sslprofile" {
  name = "System-Standard"
}

data "avi_serviceenginegroup" "se_group" {
  name      = "SE_Group_One"
  cloud_ref = data.avi_cloud.aws_cloud_cfg.id
}

resource "avi_pool" "terraform-pool-version1" {
  name                = "poolv1"
  health_monitor_refs = [data.avi_healthmonitor.system_http_healthmonitor.id]
  tenant_ref          = data.avi_tenant.default_tenant.id

  vrf_ref   = data.avi_vrfcontext.terraform_vrf.id
  cloud_ref = data.avi_cloud.aws_cloud_cfg.id

  servers {
    ip {
      type = "V4"
      addr = aws_instance.terraform-webserver[0].private_ip
    }

    availability_zone = var.aws_availability_zone

    discovered_networks {
      network_ref = "https://${data.aws_instance.avi_controller.private_ip}/api/network/${data.aws_subnet.terraform-subnets-0.id}"

      subnet {
        ip_addr {
          addr = var.aws_subnet_ip
          type = "V4"
        }

        mask = var.aws_subnet_mask
      }
    }

    hostname = aws_instance.terraform-webserver[0].private_ip
    port     = 80
  }

  servers {
    ip {
      type = "V4"
      addr = aws_instance.terraform-webserver[1].private_ip
    }

    availability_zone = var.aws_availability_zone

    discovered_networks {
      network_ref = "https://${data.aws_instance.avi_controller.private_ip}/api/network/${data.aws_subnet.terraform-subnets-0.id}"

      subnet {
        ip_addr {
          addr = var.aws_subnet_ip
          type = "V4"
        }

        mask = var.aws_subnet_mask
      }
    }

    hostname = aws_instance.terraform-webserver[1].private_ip
    port     = 80
  }

  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_pool" "terraform-pool-version2" {
  name                = "poolv2"
  health_monitor_refs = [data.avi_healthmonitor.system_http_healthmonitor.id]
  tenant_ref          = data.avi_tenant.default_tenant.id

  vrf_ref   = data.avi_vrfcontext.terraform_vrf.id
  cloud_ref = data.avi_cloud.aws_cloud_cfg.id

  servers {
    ip {
      type = "V4"
      addr = aws_instance.terraform-webserver[2].private_ip
    }

    availability_zone = var.aws_availability_zone

    discovered_networks {
      network_ref = "https://${data.aws_instance.avi_controller.private_ip}/api/network/${data.aws_subnet.terraform-subnets-0.id}"

      subnet {
        ip_addr {
          addr = var.aws_subnet_ip
          type = "V4"
        }

        mask = var.aws_subnet_mask
      }
    }

    hostname = aws_instance.terraform-webserver[2].private_ip
    port     = 80
  }

  servers {
    ip {
      type = "V4"
      addr = aws_instance.terraform-webserver[3].private_ip
    }

    availability_zone = var.aws_availability_zone

    discovered_networks {
      network_ref = "https://${data.aws_instance.avi_controller.private_ip}/api/network/${data.aws_subnet.terraform-subnets-0.id}"

      subnet {
        ip_addr {
          addr = var.aws_subnet_ip
          type = "V4"
        }

        mask = var.aws_subnet_mask
      }
    }

    hostname = aws_instance.terraform-webserver[3].private_ip
    port     = 80
  }

  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_pool" "terraform-pool-version3" {
  name                      = "poolv3"
  health_monitor_refs       = [data.avi_healthmonitor.system_http_healthmonitor.id]
  tenant_ref                = data.avi_tenant.default_tenant.id
  vrf_ref                   = data.avi_vrfcontext.terraform_vrf.id
  cloud_ref                 = data.avi_cloud.aws_cloud_cfg.id
  external_autoscale_groups = [aws_autoscaling_group.asg_based_pool.name]

  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_poolgroup" "terraform-poolgroup" {
  name       = "terraform_poolgroup"
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref  = data.avi_cloud.aws_cloud_cfg.id

  members {
    pool_ref = avi_pool.terraform-pool-version1.id
    ratio    = 100
  }

  members {
    pool_ref = avi_pool.terraform-pool-version2.id
    ratio    = 10
  }

  members {
    pool_ref = avi_pool.terraform-pool-version3.id
    ratio    = 10
  }
}

resource "avi_vsvip" "terraform-vip" {
  name            = var.avi_vip #"aws_vip"
  tenant_ref      = var.avi_username
  cloud_ref       = data.avi_cloud.aws_cloud_cfg.id
  vrf_context_ref = data.avi_vrfcontext.terraform_vrf.id

  dns_info {
    fqdn = "${var.vs_hostname}.${var.domain_name}"
  }

  vip {
    vip_id                    = "0"
    auto_allocate_ip          = true
    avi_allocated_vip         = true
    auto_allocate_floating_ip = var.floating_ip
    availability_zone         = var.aws_availability_zone
    subnet_uuid               = data.aws_subnet.terraform-subnets-0.id

    subnet {
      ip_addr {
        addr = var.aws_subnet_ip
        type = "V4"
      }

      mask = var.aws_subnet_mask
    }
  }
  vip {
    vip_id                    = "1"
    auto_allocate_ip          = true
    avi_allocated_vip         = true
    auto_allocate_floating_ip = var.floating_ip
    availability_zone         = var.aws_availability_zone
    subnet_uuid               = data.aws_subnet.terraform-subnets-1.id

    subnet {
      ip_addr {
        addr = var.aws_subnet_ip
        type = "V4"
      }

      mask = var.aws_subnet_mask
    }
  }
}

data "avi_wafpolicy" "vs_wafpolicy" {
  name = var.waf_name
}
data "avi_wafpolicy" "template_wafpolicy" {
  name = "Template-WAF-Policy-Standard"
}
resource "avi_virtualservice" "terraform-virtualservice" {
  depends_on                   = [avi_errorpageprofile.demo_errorpage]
  name                         = var.vs_name
  cloud_type                   = var.cloud_type #"CLOUD_AWS"
  cloud_ref                    = data.avi_cloud.aws_cloud_cfg.id
  pool_group_ref               = avi_poolgroup.terraform-poolgroup.id
  tenant_ref                   = data.avi_tenant.default_tenant.id
  application_profile_ref      = data.avi_applicationprofile.system_https_profile.id
  network_profile_ref          = data.avi_networkprofile.system_tcp_profile.id
  analytics_profile_ref        = data.avi_analyticsprofile.system_analytics_profile.id
  ssl_key_and_certificate_refs = [data.avi_sslkeyandcertificate.system_default_cert.id]
  ssl_profile_ref              = data.avi_sslprofile.system_standard_sslprofile.id
  se_group_ref                 = data.avi_serviceenginegroup.se_group.id
  vrf_context_ref              = data.avi_vrfcontext.terraform_vrf.id
  vsvip_ref                    = avi_vsvip.terraform-vip.id
  waf_policy_ref               = data.avi_wafpolicy.vs_wafpolicy.id == null ? data.avi_wafpolicy.template_wafpolicy.id : data.avi_wafpolicy.vs_wafpolicy.id
  error_page_profile_ref       = avi_errorpageprofile.demo_errorpage.id
  #error_page_profile_ref      = avi_errorpageprofile.cop_errorpage.id == 0 ? avi_errorpageprofile.cop_errorpage.id : ""

  services {
    enable_http2 = false
    enable_ssl   = true
    port         = 443
  }
  services {
    port = 80
    //enable_ssl     = true
    port_range_end = 80
  }
  analytics_policy {
    metrics_realtime_update {
      enabled  = true
      duration = 0
    }
  }
}

resource "aws_launch_configuration" "web_app_launch_conf" {
  name          = "${var.project_name}-${var.autoscaling_name}"
  image_id      = var.webserver_ami
  instance_type = var.webserver_instance_type
}

resource "aws_autoscaling_group" "asg_based_pool" {
  name = "${var.project_name}-${var.autoscaling_name}"
  #availability_zones   = var.aws_availability_zones
  max_size             = 2
  min_size             = 1
  force_delete         = true
  launch_configuration = aws_launch_configuration.web_app_launch_conf.name
  # If you do not use the az above you can enable the zone below
  vpc_zone_identifier = [
    data.aws_subnet.terraform-subnets-0.id,
    data.aws_subnet.terraform-subnets-1.id,
    data.aws_subnet.terraform-subnets-2.id
  ]
}


data "avi_errorpagebody" "default_error_page" {
  name = "Custom-Error-Page"
}
resource "avi_errorpageprofile" "demo_errorpage" {
  name = "Custom-Error-Page"
  error_pages {
    enable              = true
    error_page_body_ref = data.avi_errorpagebody.default_error_page.id
    index               = 0
    match {
      match_criteria = "IS_IN"
      status_codes = [
        400,
        401,
        403,
      ]
    }
  }
}
