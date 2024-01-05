provider "avi" {
  avi_username    = "admin"
  avi_tenant      = "admin"
  avi_password    = var.avi_password
  avi_controller  = var.avi_controller
  avi_version     = "30.1.2"
  avi_api_timeout = 50
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

data "avi_cloud" "default_cloud" {
  name = "Default-Cloud"
}

data "avi_serviceenginegroup" "se_group" {
  name      = "Default-Group"
  cloud_ref = data.avi_cloud.default_cloud.id
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

data "avi_vrfcontext" "global_vrf" {
  name      = "global"
  cloud_ref = data.avi_cloud.default_cloud.id
}

resource "avi_networkprofile" "test_networkprofile" {
  name       = "terraform-network-profile"
  tenant_ref = data.avi_tenant.default_tenant.id
  profile {
    type = "PROTOCOL_TYPE_TCP_PROXY"
    tcp_proxy_profile {}
  }
}

resource "avi_applicationpersistenceprofile" "test_applicationpersistenceprofile" {
  name             = "terraform-app-pers-profile"
  tenant_ref       = data.avi_tenant.default_tenant.id
  persistence_type = "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS"
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
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_virtualservice" "test_vs" {
  name           = "terraform-vs"
  pool_group_ref = avi_poolgroup.terraform-pg-1.id

  #pool_ref= "${avi_pool.testpool.id}"
  tenant_ref = data.avi_tenant.default_tenant.id
  vsvip_ref  = avi_vsvip.test_vsvip.id

  services {
    port           = 80
    enable_ssl     = true
    port_range_end = 80
  }
  cloud_type                   = "CLOUD_NONE"
  ssl_key_and_certificate_refs = [avi_sslkeyandcertificate.terraform_vs_cert.id]
  ssl_profile_ref              = data.avi_sslprofile.system_standard_sslprofile.id
}

resource "avi_healthmonitor" "test_hm_1" {
  name       = "terraform-monitor"
  type       = "HEALTH_MONITOR_HTTP"
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_pool" "terraform-pool-1" {
  name                                = "terraform-pool-1"
  health_monitor_refs                 = [avi_healthmonitor.test_hm_1.id]
  tenant_ref                          = data.avi_tenant.default_tenant.id
  cloud_ref                           = data.avi_cloud.default_cloud.id
  application_persistence_profile_ref = avi_applicationpersistenceprofile.test_applicationpersistenceprofile.id
  servers {
    ip {
      type = "V4"
      addr = var.avi_test_server_p11
    }
    port = 8080
  }
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

resource "avi_pool" "terraform-pool-2" {
  name                                = "terraform-pool-2"
  tenant_ref                          = data.avi_tenant.default_tenant.id
  cloud_ref                           = data.avi_cloud.default_cloud.id
  application_persistence_profile_ref = avi_applicationpersistenceprofile.test_applicationpersistenceprofile.id
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
  ignore_servers = true
}


resource "avi_poolgroup" "terraform-pg-1" {
  name       = "terraform-pg-1"
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref  = data.avi_cloud.default_cloud.id
  members {
    pool_ref         = avi_pool.terraform-pool-1.id
    ratio            = 100
    deployment_state = "IN_SERVICE"
  }
  members {
    pool_ref         = avi_pool.terraform-pool-2.id
    ratio            = 0
    deployment_state = "OUT_OF_SERVICE"
  }
}

resource "avi_server" "test_server_p21" {
  ip       = var.avi_test_server_p21
  port     = "80"
  pool_ref = avi_pool.terraform-pool-2.id
  hostname = "foo"
}

resource "avi_server" "test_server_p22" {
  ip       = var.avi_test_server_p22
  port     = "80"
  pool_ref = avi_pool.terraform-pool-2.id
  hostname = "bar1"
}

#resource "avi_server" "test_server_p23" {
#ip       = var.avi_test_server_p23
#  port     = "80"
#  pool_ref = avi_pool.terraform-pool-2.id
#  hostname = "goo"
#}

resource "avi_server" "test_server" {
  ip       = var.avi_test_server_p12
  port     = "80"
  pool_ref = avi_pool.terraform-pool-1.id
  hostname = var.avi_test_server_p12
}

resource "avi_sslkeyandcertificate" "terraform_vs_cert" {
  name = "terraform_vs_cert"
  key  = file("${path.module}/app_cert.key")
  certificate {
    self_signed = true
    certificate = file("${path.module}/app_cert.crt")
  }
  type = "SSL_CERTIFICATE_TYPE_VIRTUALSERVICE"
}
