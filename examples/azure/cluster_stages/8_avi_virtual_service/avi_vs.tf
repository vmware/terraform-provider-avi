
resource "avi_vsvip" "azure-vs-vsvip" {
  name            = "azure_vip"
  tenant_ref      = data.avi_tenant.default_tenant.id
  cloud_ref       = data.avi_cloud.azure_cloud_cfg.id
  vrf_context_ref = data.avi_vrfcontext.terraform_vrf.id

  vip {
    vip_id                    = "0"
    auto_allocate_ip          = true
    avi_allocated_vip         = true
    avi_allocated_fip         = false
    subnet_uuid               = data.azurerm_subnet.terraform_subnet.name
    auto_allocate_floating_ip = var.floating_ip
    enabled                   = true
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
  depends_on                   = [avi_errorpageprofile.demo_errorpage]
  name                         = var.vs_name
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
  waf_policy_ref               = data.avi_wafpolicy.vs_wafpolicy.id == null ? data.avi_wafpolicy.template_wafpolicy.id : data.avi_wafpolicy.vs_wafpolicy.id
  vsvip_ref = avi_vsvip.azure-vs-vsvip.id
  error_page_profile_ref       = avi_errorpageprofile.demo_errorpage.id

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
