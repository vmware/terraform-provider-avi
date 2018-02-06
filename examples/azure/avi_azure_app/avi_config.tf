provider "azurerm" {
  subscription_id 	= "${var.subscription_id}"
  client_id 		= "${var.client_id}"
  client_secret 	= "${var.client_secret}"
  tenant_id 		= "${var.tenant_id}"
}

data "azurerm_public_ip" "terraform_controller" {
  name = "Terraform-PIP"
  resource_group_name = "${var.resource_group_name}"
}

data "azurerm_public_ip" "terraform_webserver1" {
  name = "Terraform-PIP-WebVM1"
  resource_group_name = "${var.resource_group_name}"
}

data "azurerm_public_ip" "terraform_webserver2" {
  name = "Terraform-PIP-WebVM2"
  resource_group_name = "${var.resource_group_name}"
}

data "azurerm_public_ip" "terraform_webserver3" {
  name = "Terraform-PIP-WebVM3"
  resource_group_name = "${var.resource_group_name}"
}

data "azurerm_public_ip" "terraform_webserver4" {
  name = "Terraform-PIP-WebVM4"
  resource_group_name = "${var.resource_group_name}"
}

data "azurerm_virtual_network" "terraform_vnet" {
  name 			= "Terraform-VNet"
  resource_group_name   = "${var.resource_group_name}"
}

output "subnets" {
  value = "${data.azurerm_virtual_network.terraform_vnet.subnets}"
}

provider "avi" {
  avi_username   = "${var.avi_username}"
  avi_password   = "${var.avi_password}"
  avi_controller = "${data.azurerm_public_ip.terraform_controller.ip_address}"
  avi_tenant     = "admin"
}

data "avi_applicationprofile" "system_http_profile" {
  name = "System-HTTP"
}

data "avi_applicationprofile" "system_https_profile" {
  name = "System-Secure-HTTP"
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

data "avi_cloud" "azure_cloud_cfg" {
  name = "azure_cloud_cfg"
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = "${data.avi_cloud.azure_cloud_cfg.id}"
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
  name      = "Default-Group"
  cloud_ref = "${data.avi_cloud.azure_cloud_cfg.id}"
}

resource "avi_pool" "terraform-pool-version1" {
  name                = "poolv1"
  health_monitor_refs = ["${data.avi_healthmonitor.system_http_healthmonitor.id}"]
  server_count        = 2
  tenant_ref          = "${data.avi_tenant.default_tenant.id}"

  vrf_ref   = "${data.avi_vrfcontext.terraform_vrf.id}"
  cloud_ref = "${data.avi_cloud.azure_cloud_cfg.id}"

  servers {
    ip = {
      type = "V4"
      addr = "${data.azurerm_public_ip.terraform_webserver1.ip_address}"
    }

    discovered_networks = {
      network_ref = "https://${data.azurerm_public_ip.terraform_controller.ip_address}/api/network/${data.azurerm_virtual_network.terraform_vnet.subnets[0]}"

      subnet = {
        ip_addr = {
          addr = "${var.azure_subnet_ip}"
          type = "V4"
        }

        mask = "${var.azure_subnet_mask}"
      }
    }

    hostname = "${data.azurerm_public_ip.terraform_webserver1.ip_address}"
    port     = 80
  }

  servers {
    ip = {
      type = "V4"
      addr = "${data.azurerm_public_ip.terraform_webserver2.ip_address}"
    }

    discovered_networks = {
      network_ref = "https://${data.azurerm_public_ip.terraform_controller.ip_address}/api/network/${data.azurerm_virtual_network.terraform_vnet.subnets[0]}"

      subnet = {
        ip_addr = {
          addr = "${var.azure_subnet_ip}"
          type = "V4"
        }

        mask = "${var.azure_subnet_mask}"
      }
    }

    hostname = "${data.azurerm_public_ip.terraform_webserver2.ip_address}"
    port     = 80
  }

  fail_action = {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_pool" "terraform-pool-version2" {
  name                = "poolv2"
  health_monitor_refs = ["${data.avi_healthmonitor.system_http_healthmonitor.id}"]
  server_count        = 2
  tenant_ref          = "${data.avi_tenant.default_tenant.id}"

  vrf_ref   = "${data.avi_vrfcontext.terraform_vrf.id}"
  cloud_ref = "${data.avi_cloud.azure_cloud_cfg.id}"

  servers {
    ip = {
      type = "V4"
      addr = "${data.azurerm_public_ip.terraform_webserver3.ip_address}"
    }

    discovered_networks = {
      network_ref = "https://${data.azurerm_public_ip.terraform_controller.ip_address}/api/network/${data.azurerm_virtual_network.terraform_vnet.subnets[0]}"

      subnet = {
        ip_addr = {
          addr = "${var.azure_subnet_ip}"
          type = "V4"
        }

        mask = "${var.azure_subnet_mask}"
      }
    }

    hostname = "${data.azurerm_public_ip.terraform_webserver3.ip_address}"
    port     = 80
  }

  servers {
    ip = {
      type = "V4"
      addr = "${data.azurerm_public_ip.terraform_webserver4.ip_address}"
    }

    discovered_networks = {
      network_ref = "https://${data.azurerm_public_ip.terraform_controller.ip_address}/api/network/${data.azurerm_virtual_network.terraform_vnet.subnets[0]}"

      subnet = {
        ip_addr = {
          addr = "${var.azure_subnet_ip}"
          type = "V4"
        }

        mask = "${var.azure_subnet_mask}"
      }
    }

    hostname = "${data.azurerm_public_ip.terraform_webserver4.ip_address}"
    port     = 80
  }

  fail_action = {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_poolgroup" "terraform-poolgroup" {
  name       = "terraform_poolgroup"
  tenant_ref = "${data.avi_tenant.default_tenant.id}"
  cloud_ref  = "${data.avi_cloud.azure_cloud_cfg.id}"

  members = {
    pool_ref = "${avi_pool.terraform-pool-version1.id}"
    ratio    = 100
  }

  members = {
    pool_ref = "${avi_pool.terraform-pool-version2.id}"
    ratio    = 0
  }
}

resource "avi_vsvip" "terraform-vip" {
  name            = "azure_vip"
  tenant_ref      = "${data.avi_tenant.default_tenant.id}"
  cloud_ref       = "${data.avi_cloud.azure_cloud_cfg.id}"
  vrf_context_ref = "${data.avi_vrfcontext.terraform_vrf.id}"

  vip {
    auto_allocate_ip  = true
    avi_allocated_vip = true
    avi_allocated_fip = true
    auto_allocate_floating_ip = true
    subnet_uuid       = "${data.azurerm_virtual_network.terraform_vnet.subnets[0]}"

    subnet = {
      ip_addr = {
        addr = "${var.azure_subnet_ip}"
        type = "V4"
      }

      mask = "${var.azure_subnet_mask}"
    }
  }
}

resource "avi_virtualservice" "terraform-virtualservice" {
  name                         = "azure_vs"
  pool_group_ref               = "${avi_poolgroup.terraform-poolgroup.id}"
  tenant_ref                   = "${data.avi_tenant.default_tenant.id}"
  cloud_type                   = "CLOUD_AZURE"
  application_profile_ref      = "${data.avi_applicationprofile.system_https_profile.id}"
  network_profile_ref          = "${data.avi_networkprofile.system_tcp_profile.id}"
  cloud_ref                    = "${data.avi_cloud.azure_cloud_cfg.id}"
  analytics_profile_ref        = "${data.avi_analyticsprofile.system_analytics_profile.id}"
  ssl_key_and_certificate_refs = ["${data.avi_sslkeyandcertificate.system_default_cert.id}"]
  ssl_profile_ref              = "${data.avi_sslprofile.system_standard_sslprofile.id}"
  se_group_ref                 = "${data.avi_serviceenginegroup.se_group.id}"
  vrf_context_ref              = "${data.avi_vrfcontext.terraform_vrf.id}"

  //vsvip_ref                    = "${avi_vsvip.terraform-vip.id}"

  vip {
    auto_allocate_ip  = true
    avi_allocated_vip = true
    avi_allocated_fip = true
    auto_allocate_floating_ip = true
    subnet_uuid       = "${data.azurerm_virtual_network.terraform_vnet.subnets[0]}"

    subnet = {
      ip_addr = {
        addr = "${var.azure_subnet_ip}"
        type = "V4"
      }

      mask = "${var.azure_subnet_mask}"
    }
  }
  services {
    port           = 80
    enable_ssl     = true
    port_range_end = 80
  }
  analytics_policy {
    metrics_realtime_update = {
      enabled  = true
      duration = 0
    }
  }
}
