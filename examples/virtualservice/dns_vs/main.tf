provider "avi" {
  avi_username   = var.avi_username
  avi_tenant     = "admin"
  avi_password   = var.avi_password
  avi_controller = var.avi_controller
  avi_version    = var.avi_version
  avi_api_timeout    = 50
}

data "avi_tenant" "default_tenant" {
  name = "admin"
}
resource "avi_cloud" "cloud_none" {
  vtype                        = "CLOUD_NONE"
  license_tier                 = "ENTERPRISE"
  dhcp_enabled                 = false
  prefer_static_routes         = false
  license_type                 = "LIC_CORES"
  mtu                          = 1500
  apic_mode                    = false
  state_based_dns_registration = true
  name                         = "test-tf-cloud"
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_sslkeyandcertificate" "ssl_cert1" {
  name = "terraform_vs_cert"
  key = file("${path.module}/../../basic/app_cert.key")
  certificate {
    self_signed = true
    certificate = file("${path.module}/../../basic/app_cert.crt")
  }
  type= "SSL_CERTIFICATE_TYPE_VIRTUALSERVICE"
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_sslprofile" "ssl_profile1" {
  name = "tf-ssl-profile"
  accepted_versions {
    type = "SSL_VERSION_TLS1"
  }
  accepted_versions {
      type = "SSL_VERSION_TLS1_1"
    }
  accepted_versions {
    type = "SSL_VERSION_TLS1_2"
  }
  cipher_enums = [
    "TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
    "TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
    "TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
    "TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
  ]
  type = "SSL_PROFILE_TYPE_APPLICATION"
}
resource "avi_applicationprofile" "application_profile1" {
  name = "terraform_dns_application_profile"
  type = "APPLICATION_PROFILE_TYPE_DNS"
  dns_service_profile {
    num_dns_ip = 1
    ttl = 30
    error_response = "DNS_ERROR_RESPONSE_NONE"
    edns = true
    edns_client_subnet_prefix_len = 24
    dns_over_tcp_enabled = true
    aaaa_empty_response = true
    negative_caching_ttl = 30
    admin_email = "hostmaster"
    ecs_stripping_enabled = true
  }
  preserve_client_ip = false
  preserve_client_port = false
  preserve_dest_ip_port = false
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_networkprofile" "network_profile1" {
  name = "tf-udp-pkt-network-profile"
  profile {
    type = "PROTOCOL_TYPE_UDP_FAST_PATH"
    udp_fast_path_profile {
      session_idle_timeout = 10
      per_pkt_loadbalance = true
      snat = true
    }
  }
  connection_mirror = false
  tenant_ref = data.avi_tenant.default_tenant.id
}


resource "avi_pool" "lb_pool" {
  name         = var.pool_name
  lb_algorithm = var.lb_algorithm
  servers {
    ip {
          type = "V4"
      addr = var.server1_ip
    }
    port = var.server1_port
  }
  cloud_ref = avi_cloud.cloud_none.id
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_poolgroup" "poolgroup1" {
  name         = var.poolgroup_name
  members {
        pool_ref = avi_pool.lb_pool.id
        ratio = 100
  }
  cloud_ref = avi_cloud.cloud_none.id
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_vsvip" "test_vsvip" {
  name = "terraform-vip"
  vip {
    vip_id = "0"
    ip_address {
      type = "V4"
      addr = var.avi_terraform_vs_vip
    }
  }
  cloud_ref = avi_cloud.cloud_none.id
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_virtualservice" "dns_vs" {
  name                          = var.vs_name
  pool_group_ref                = avi_poolgroup.poolgroup1.id
  tenant_ref                    = data.avi_tenant.default_tenant.id
  vsvip_ref                     = avi_vsvip.test_vsvip.id
  cloud_ref                     = avi_cloud.cloud_none.id
  ssl_key_and_certificate_refs  = [avi_sslkeyandcertificate.ssl_cert1.id]
  ssl_profile_ref               = avi_sslprofile.ssl_profile1.id
  application_profile_ref       = avi_applicationprofile.application_profile1.id
  network_profile_ref           = avi_networkprofile.network_profile1.id
  services {
    port           = var.vs_port
    enable_ssl     = true
  }
}
