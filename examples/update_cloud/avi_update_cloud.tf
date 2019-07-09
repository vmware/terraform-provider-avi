provider "avi" {
  avi_username   = var.avi_username
  avi_tenant     = var.avi_tenant
  avi_password   = var.avi_password
  avi_controller = var.avi_controller
}

resource "avi_serviceenginegroup" "test_se" {
  name                = "Default-Group"
  license_type        = "LIC_SE_BANDWIDTH"
  se_bandwidth_type   = "SE_BANDWIDTH_200M"
  algo                = "PLACEMENT_ALGO_PACKED"
  min_scaleout_per_vs = 2
  ha_mode             = "HA_MODE_SHARED_PAIR"
}

resource "avi_tenant" "test_tenant" {
  name = "test-tenant"
}

resource "avi_tenant" "avi_tenant" {
  name = "avinetworks"
}

resource "avi_sslprofile" "system_standard_sslprofile" {
  name = "System-Standard"
  accepted_versions {
    type = "SSL_VERSION_TLS1"
  }
  accepted_versions {
    type = "SSL_VERSION_TLS1_2"
  }
  cipher_enums = [
    "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA384",
    "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
    "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
    "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA",
    "TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA256",
    "TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA",
  ]
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}

resource "avi_role" "testRole" {
  privileges {
    resource = "PERMISSION_CONTROLLER"
    type     = "NO_ACCESS"
  }
  privileges {
    resource = "PERMISSION_VIRTUALSERVICE"
    type     = "READ_ACCESS"
  }
  privileges {
    resource = "PERMISSION_POOL"
    type     = "READ_ACCESS"
  }
  privileges {
    resource = "PERMISSION_POOLGROUP"
    type     = "READ_ACCESS"
  }
  tenant_ref = data.avi_tenant.default_tenant.id
  name       = "testApplication-Admin"
}

resource "avi_alertemailconfig" "testalertemailconfig" {
  name        = "aec-abc"
  tenant_ref  = data.avi_tenant.default_tenant.id
  cc_emails   = "admin@avicontroller.net"
  description = "test alert email"
  to_emails   = "admin@avicontroller.net"
}

resource "avi_alertsyslogconfig" "testalertsyslogconfig" {
  name        = "asyc-test"
  description = "test alert syslog"
  tenant_ref  = data.avi_tenant.default_tenant.id
}

resource "avi_cloudconnectoruser" "custom_cloud_ref" {
  azure_userpass {
    password    = var.azure_password
    username    = var.azure_username
    tenant_name = var.azure_tenant_name
  }
  name = "test_user"
}

resource "avi_cloud" "azurcloud" {
  vtype                        = "CLOUD_AZURE"
  license_tier                 = "ENTERPRISE_18"
  dhcp_enabled                 = true
  prefer_static_routes         = false
  license_type                 = "LIC_CORES"
  mtu                          = 1500
  apic_mode                    = false
  state_based_dns_registration = true
  name                         = "Default-Cloud"
  azure_configuration {
    location              = "westus"
    resource_group        = var.resource_group
    subscription_id       = var.subscription_id
    use_azure_dns         = "true"
    cloud_credentials_ref = avi_cloudconnectoruser.custom_cloud_ref.id
    network_info {
      se_network_id      = "default"
      virtual_network_id = var.virtual_network_id
    }
  }
}

