provider "azurerm" {
  subscription_id = var.subscription_id
  client_id       = var.client_id
  client_secret   = var.client_secret
  tenant_id       = var.tenant_id
}

data "azurerm_network_interface" "controller_nic" {
  name                = "${var.project_name}-terraform-network-interface"
  resource_group_name = "${var.project_name}-terraform-resource-group"
  //resource_group_name = "${var.resource_group_name}"
}

data "azurerm_subnet" "terraform_subnet" {
  name                 = "${var.project_name}-terraform-subnet"
  virtual_network_name = "${var.project_name}-terraform-virtual-network"
  resource_group_name  = "${var.project_name}-terraform-resource-group"
}

provider "avi" {
  avi_username   = var.avi_username
  avi_password   = var.avi_password
  avi_controller = data.azurerm_network_interface.controller_nic.private_ip_address
  avi_tenant     = "admin"
  avi_version    = var.avi_version
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
  name = "AZURE"
}

data "avi_vrfcontext" "terraform_vrf" {
  name      = "global"
  cloud_ref = data.avi_cloud.azure_cloud_cfg.id
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
  cloud_ref = data.avi_cloud.azure_cloud_cfg.id
}

data "avi_poolgroup" "azure-poolgroup" {
  name       = "azure_poolgroup"
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref  = data.avi_cloud.azure_cloud_cfg.id
}

resource "avi_vsvip" "azure-vs-vsvip" {
  name            = "azure_vip"
  tenant_ref      = data.avi_tenant.default_tenant.id
  cloud_ref       = data.avi_cloud.azure_cloud_cfg.id
  vrf_context_ref = data.avi_vrfcontext.terraform_vrf.id

  vip {
    vip_id            = "0"
    auto_allocate_ip  = true
    avi_allocated_vip = true
    avi_allocated_fip = false
    subnet_uuid       = data.azurerm_subnet.terraform_subnet.name
    enabled           = true
    subnet {
      ip_addr {
        addr = var.azure_vip_subnet_ip
        type = "V4"
      }

      mask = var.azure_vip_subnet_mask
    }
  }
}

resource "avi_virtualservice" "azure-virtualservice" {
  name                         = "azure_vs"
  pool_group_ref               = data.avi_poolgroup.azure-poolgroup.id
  tenant_ref                   = data.avi_tenant.default_tenant.id
  cloud_type                   = "CLOUD_AZURE"
  application_profile_ref      = data.avi_applicationprofile.system_https_profile.id
  network_profile_ref          = data.avi_networkprofile.system_tcp_profile.id
  cloud_ref                    = data.avi_cloud.azure_cloud_cfg.id
  analytics_profile_ref        = data.avi_analyticsprofile.system_analytics_profile.id
  ssl_key_and_certificate_refs = [data.avi_sslkeyandcertificate.system_default_cert.id]
  ssl_profile_ref              = data.avi_sslprofile.system_standard_sslprofile.id
  se_group_ref                 = data.avi_serviceenginegroup.se_group.id
  vrf_context_ref              = data.avi_vrfcontext.terraform_vrf.id
  scaleout_ecmp                = true
  enabled                      = true

  vsvip_ref                    = avi_vsvip.azure-vs-vsvip.id

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
    metrics_realtime_update {
      enabled  = true
      duration = 0
    }
  }
}

