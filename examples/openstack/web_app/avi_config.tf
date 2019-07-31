provider "openstack" {
  user_name   = var.openstack_username
  tenant_name = var.openstack_tenant_name
  password    = var.openstack_password
  auth_url    = var.openstack_url
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = var.avi_controller
  avi_tenant     = var.avi_tenant
}

data "openstack_networking_network_v2" "network" {
  name = var.network_name
}

data "openstack_networking_subnet_v2" "terraform-subnets-0" {
  name = var.openstack_subnet
}

data "avi_applicationprofile" "system_http_profile" {
  name = "System-HTTP"
}

data "avi_applicationprofile" "system_https_profile" {
  name = "System-Secure-HTTP"
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

data "avi_healthmonitor" "system_http_healthmonitor" {
  name = "System-HTTP"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

data "avi_cloud" "op_cloud_cfg" {
  name = "OpenStack"
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = data.avi_cloud.op_cloud_cfg.id
}

data "avi_serviceenginegroup" "se_group" {
  name      = "Default-Group"
  cloud_ref = data.avi_cloud.op_cloud_cfg.id
}

resource "avi_pool" "terraform-pool-version1" {
  name                = "poolv1"
  health_monitor_refs = [data.avi_healthmonitor.system_http_healthmonitor.id]
  tenant_ref          = data.avi_tenant.default_tenant.id

  vrf_ref   = data.avi_vrfcontext.terraform_vrf.id
  cloud_ref = data.avi_cloud.op_cloud_cfg.id

  servers {
    ip {
      type = "V4"
      addr = var.avi_pool_server_1
    }

    discovered_networks {
      network_ref = "https://${var.avi_controller}/api/network/${data.openstack_networking_subnet_v2.terraform-subnets-0.id}"

      subnet {
        ip_addr {
          addr = var.op_subnet_ip
          type = "V4"
        }

        mask = var.op_subnet_mask
      }
    }
    port     = 80
    hostname = "pool-server-1"
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
  cloud_ref = data.avi_cloud.op_cloud_cfg.id

  servers {
    ip {
      type = "V4"
      addr = var.avi_pool_server_2
    }

    discovered_networks {
      network_ref = "https://${var.avi_controller}/api/network/${data.openstack_networking_subnet_v2.terraform-subnets-0.id}"

      subnet {
        ip_addr {
          addr = var.op_subnet_ip
          type = "V4"
        }

        mask = var.op_subnet_mask
      }
    }
    port     = 80
    hostname = "pool-server-2"
  }

  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_poolgroup" "terraform-poolgroup" {
  name       = "terraform_poolgroup"
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref  = data.avi_cloud.op_cloud_cfg.id

  members {
    pool_ref = avi_pool.terraform-pool-version1.id
    ratio    = 100
  }

  members {
    pool_ref = avi_pool.terraform-pool-version2.id
    ratio    = 100
  }
}

resource "avi_vsvip" "terraform-vip" {
  name = "aws_vip"
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref = data.avi_cloud.op_cloud_cfg.id
  vrf_context_ref = data.avi_vrfcontext.terraform_vrf.id

  vip {
    vip_id            = "0"
    auto_allocate_ip  = true
    avi_allocated_vip = true
    subnet_uuid       = data.openstack_networking_subnet_v2.terraform-subnets-0.id

    subnet {
      ip_addr {
        addr = var.op_subnet_ip
        type = "V4"
      }

      mask = var.op_subnet_mask
    }
  }
}
resource "avi_virtualservice" "terraform-virtualservice" {
  name                         = "op_vs"
  cloud_type                   = "CLOUD_OPENSTACK"
  cloud_ref                    = data.avi_cloud.op_cloud_cfg.id
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

  services {
    port           = 80
    port_range_end = 80
  }
  analytics_policy {
    metrics_realtime_update {
      enabled  = true
      duration = 0
    }
  }
}

