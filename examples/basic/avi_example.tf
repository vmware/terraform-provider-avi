provider "avi" {
  avi_username = "admin"
  avi_tenant = "admin"
  avi_password = "${var.avi_password}"
  avi_controller= "10.10.25.42"
}

data "avi_applicationprofile" "system_http_profile" {
  name= "System-HTTP"
}

data "avi_applicationprofile" "system_https_profile" {
  name= "System-Secure-HTTP"
}

data "avi_tenant" "default_tenant" {
  name= "admin"
}

data "avi_cloud" "default_cloud" {
  name= "Default-Cloud"
}

data "avi_serviceenginegroup" "se_group" {
  name = "Default-Group"
  cloud_ref = "${data.avi_cloud.default_cloud.id}"
}

data "avi_networkprofile" "system_tcp_profile" {
  name= "System-TCP-Proxy"
}

data "avi_analyticsprofile" "system_analytics_profile" {
  name= "System-Analytics-Profile"
}

data "avi_sslkeyandcertificate" "system_default_cert" {
  name= "System-Default-Cert"
}

data "avi_sslprofile" "system_standard_sslprofile" {
  name= "System-Standard"
}

data "avi_vrfcontext" "global_vrf" {
  name= "global"
  cloud_ref = "${data.avi_cloud.default_cloud.id}"
}




resource "avi_networkprofile" "test_networkprofile" {
  name= "networkprofile-42"
  tenant_ref= "${data.avi_tenant.default_tenant.id}"
  profile{
    type= "PROTOCOL_TYPE_TCP_PROXY"
  }
}

resource "avi_applicationpersistenceprofile" "test_applicationpersistenceprofile" {
  name = "applicationpersistence-42"
  tenant_ref= "${data.avi_tenant.default_tenant.id}"
  persistence_type = "PERSISTENCE_TYPE_CLIENT_IP_ADDRESS"
}

resource "avi_vsvip" "test_vsvip" {
  name= "vip-42"
  vip {
    vip_id= "0"
    ip_address {
      type= "V4",
      addr= "10.90.64.88",
    }
  }
}

resource "avi_virtualservice" "test_vs" {
  name= "vs-42"
  pool_ref= "${avi_pool.testpool.id}"
  tenant_ref= "${data.avi_tenant.default_tenant.id}"
  vsvip_ref = "${avi_vsvip.test_vsvip.id}"
  vip {
    vip_id= "0"
    ip_address {
      type= "V4",
      addr= "10.90.64.88",
    }
  }
  services {
    port= 80
    enable_ssl= true
    port_range_end= 80
  }
  cloud_type = "CLOUD_VCENTER"
  ssl_key_and_certificate_refs= ["${data.avi_sslkeyandcertificate.system_default_cert.id}"]
  ssl_profile_ref= "${data.avi_sslprofile.system_standard_sslprofile.id}"
}


resource "avi_healthmonitor" "test_hm_1" {
  name = "healthmonitor-42"
  type = "HEALTH_MONITOR_HTTP"
}


resource "avi_pool" "testpool" {
  name= "pool-42",
  health_monitor_refs= ["${avi_healthmonitor.test_hm_1.id}"]
  tenant_ref= "${data.avi_tenant.default_tenant.id}"
  cloud_ref= "${data.avi_cloud.default_cloud.id}"
  application_persistence_profile_ref= "${avi_applicationpersistenceprofile.test_applicationpersistenceprofile.id}"
  servers {
    ip= {
      type= "V4",
      addr= "10.90.64.66",
    }
    port= 8080
  }
  fail_action= {
    type= "FAIL_ACTION_CLOSE_CONN"
  }
}
