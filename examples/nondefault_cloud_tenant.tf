provider "avi" {
  avi_username   = "admin"
  avi_tenant     = "admin"
  avi_password   = "something"
  avi_controller = "10.10.25.42"
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
  name = "Default-Group"
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
  name = "global"
}

resource "avi_tenant" "test_tenant" {
  name = "tenant1"
}

resource "avi_cloud" "test_cloud" {
  name       = "Cloud1"
  vtype      = "CLOUD_NONE"
  tenant_ref = avi_tenant.test_tenant.id
}

resource "avi_networkprofile" "test_networkprofile" {
  name       = "networkprofile1"
  tenant_ref = avi_tenant.test_tenant.id
  profile {
    type = "PROTOCOL_TYPE_TCP_PROXY"
  }
}

resource "avi_applicationpersistenceprofile" "test_applicationpersistenceprofile" {
  name             = "applicationpersistence1"
  tenant_ref       = avi_tenant.test_tenant.id
  persistence_type = "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS"
}

resource "avi_vsvip" "test_vsvip" {
  name = "vip1"
  vip {
    vip_id = "0"
    ip_address {
      type = "V4"
      addr = "10.90.64.88"
    }
  }
}

resource "avi_virtualservice" "test_vs" {
  name       = "vs1"
  pool_ref   = avi_pool.testpool.id
  tenant_ref = avi_tenant.test_tenant.id
  cloud_ref  = avi_cloud.test_cloud.id

  #application_profile_ref= "${data.avi_applicationprofile.system_https_profile.id}"
  network_profile_ref = avi_networkprofile.test_networkprofile.id
  vsvip_ref           = avi_vsvip.test_vsvip.id
  vip {
    vip_id = "0"
    ip_address {
      type = "V4"
      addr = "10.90.64.88"
    }
  }
  services {
    port           = 80
    enable_ssl     = true
    port_range_end = 80
  }
  #se_group_ref= "${data.avi_serviceenginegroup.se_group.id}"
  #analytics_profile_ref= "${data.avi_analyticsprofile.system_analytics_profile.id}"
  #ssl_key_and_certificate_refs= ["${data.avi_sslkeyandcertificate.system_default_cert.id}"]
  #ssl_profile_ref= "${data.avi_sslprofile.system_standard_sslprofile.id}"
  # vrf_context_ref= "${data.avi_vrfcontext.global_vrf.id}"
}

resource "avi_healthmonitor" "test_hm_1" {
  name       = "healthmonitor1"
  type       = "HEALTH_MONITOR_HTTP"
  tenant_ref = avi_tenant.test_tenant.id
}

resource "avi_pool" "testpool" {
  name                                = "pool1"
  health_monitor_refs                 = [avi_healthmonitor.test_hm_1.id]
  tenant_ref                          = avi_tenant.test_tenant.id
  cloud_ref                           = avi_cloud.test_cloud.id
  application_persistence_profile_ref = avi_applicationpersistenceprofile.test_applicationpersistenceprofile.id
  servers {
    ip {
      type = "V4"
      addr = "10.90.64.66"
    }
    port = 8080
  }
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
}

