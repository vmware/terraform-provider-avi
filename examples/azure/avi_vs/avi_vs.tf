provider "azurerm" {
  subscription_id 	= "${var.subscription_id}"
  client_id 		= "${var.client_id}"
  client_secret 	= "${var.client_secret}"
  tenant_id 		= "${var.tenant_id}"
}


data "azurerm_public_ip" "terraform_controller" {
  name = "${var.avi_controller_name}"
  resource_group_name = "${var.resource_group_name}"
}


output "controlller_ip" {
  value = "${data.azurerm_public_ip.terraform_controller.ip_address}"
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
  name = "Default-Cloud"
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

data "avi_pool" "azure-pool-v1" {
  name = "azure_poolv1"
  tenant_ref = "${data.avi_tenant.default_tenant.id}"
  cloud_ref  = "${data.avi_cloud.azure_cloud_cfg.id}"
}

data "avi_pool" "azure-pool-v2" {
  name = "azure_poolv2"
  tenant_ref = "${data.avi_tenant.default_tenant.id}"
  cloud_ref  = "${data.avi_cloud.azure_cloud_cfg.id}"
}

resource "azurerm_subnet" "terraform_vip_subnet" {
  name                 = "${var.project_name}_terraform_vip_subnet"
  resource_group_name  = "${var.resource_group_name}"
  virtual_network_name = "${var.azure_vnet}"
  #virtual_network_name = "${var.project_name}"
  address_prefix       = "${var.azure_vip_subnet_ip}/24"
}



resource "avi_poolgroup" "azure-poolgroup" {
  name       = "azure_poolgroup"
  tenant_ref = "${data.avi_tenant.default_tenant.id}"
  cloud_ref  = "${data.avi_cloud.azure_cloud_cfg.id}"

  members = {
    pool_ref = "${data.avi_pool.azure-pool-v1.id}"
    ratio    = 100
  }

  members = {
    pool_ref = "${data.avi_pool.azure-pool-v2.id}"
    ratio    = 50
  }

}

resource "avi_virtualservice" "azure-virtualservice" {
  name                         = "azure_vs"
  pool_group_ref               = "${avi_poolgroup.azure-poolgroup.id}"
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
  scaleout_ecmp = true
  enabled = true
  vip {
    auto_allocate_ip  = true
    avi_allocated_vip = true
    avi_allocated_fip = false
    # auto_allocate_floating_ip = true
    subnet_uuid       = "${azurerm_subnet.terraform_vip_subnet.name}"

    subnet = {
      ip_addr = {
        addr = "${var.azure_vip_subnet_ip}"
        type = "V4"
      }

      mask = "${var.azure_vip_subnet_mask}"
    }
  }
  services {
    port           = 80
    enable_ssl     = true
    port_range_end = 80
  }
  services {
    port           = 443
    enable_ssl     = true
    port_range_end = 443
  }
  analytics_policy {
    metrics_realtime_update = {
      enabled  = true
      duration = 0
    }
  }
}
