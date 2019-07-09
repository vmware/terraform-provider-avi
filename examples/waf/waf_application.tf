provider "avi" {
  avi_username   = "admin"
  avi_tenant     = "admin"
  avi_password   = var.avi_password
  avi_controller = var.avi_controller
  avi_version    = "18.2.5"
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

data "avi_wafpolicy" "waf_app_learning_policy" {
  name = "waf_app_learning_policy"
}

resource "avi_networkprofile" "test_networkprofile" {
  name       = "terraform-network-profile"
  tenant_ref = data.avi_tenant.default_tenant.id
  profile {
    type = "PROTOCOL_TYPE_TCP_PROXY"
  }
}

// WAF enabled application
resource "avi_virtualservice" "waf_vs" {
  name           = "waf_vs"
  pool_group_ref = avi_poolgroup.waf_app_pg.id
  tenant_ref     = data.avi_tenant.default_tenant.id
  vsvip_ref      = avi_vsvip.vip_waf_app.id
  services {
    port           = 443
    enable_ssl     = true
    port_range_end = 443
  }
  application_profile_ref      = data.avi_applicationprofile.system_https_profile.id
  cloud_type                   = "CLOUD_VCENTER"
  ssl_key_and_certificate_refs = [data.avi_sslkeyandcertificate.system_default_cert.id]
  ssl_profile_ref              = data.avi_sslprofile.system_standard_sslprofile.id
  waf_policy_ref               = data.avi_wafpolicy.waf_app_learning_policy.id
  analytics_policy {
    metrics_realtime_update {
      enabled  = true
      duration = 0
    }
  }
}

resource "avi_vsvip" "vip_waf_app" {
  name = "vip_waf_app"
  vip {
    vip_id = "1"
    ip_address {
      type = "V4"
      addr = "10.90.64.243"
    }
  }
  tenant_ref = data.avi_tenant.default_tenant.id
}

/*// WAF profile with learning enabled
resource "avi_wafprofile" "waf_learning_profile1" {
  name= "waf_learning_profile1"
  config {
    learning_params {
      enable_per_uri_learning = true
      max_uris = 100
      min_hits_to_learn = 20
      sampling_percent = 100
      update_interval = 1
    }
    max_execution_time = 50
    min_confidence = "CONFIDENCE_VERY_HIGH"
    enable_auto_rule_updates = true
  }
  files {
    data= "p1a1\np1a2\np1a3",
    name= "file1"
  }
  files {
    data= "p1b1\np1b2\np1b3",
    name= "file2"
  }
  config {
    client_request_max_body_size= 1024
  }
}*/

resource "avi_healthmonitor" "test_hm_1" {
  name       = "terraform-monitor"
  type       = "HEALTH_MONITOR_HTTP"
  tenant_ref = data.avi_tenant.default_tenant.id
}

// Pool Group to allow traffic engineering across different versions of application
resource "avi_poolgroup" "waf_app_pg" {
  name       = "waf_app_pg"
  tenant_ref = data.avi_tenant.default_tenant.id
  members {
    pool_ref = avi_pool.waf_pool_v1.id
    ratio    = 100
  }
  /*
  members = {
    pool_ref = "${avi_pool.waf_pool_v2.id}"
    ratio    = 100
  }*/
}

// autoscaling enabled pool of application instances with version 1
resource "avi_pool" "waf_pool_v1" {
  name                = "waf_pool_v1"
  health_monitor_refs = [avi_healthmonitor.test_hm_1.id]
  tenant_ref          = data.avi_tenant.default_tenant.id
  cloud_ref           = data.avi_cloud.default_cloud.id
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
  ignore_servers = true
  analytics_policy {
    enable_realtime_metrics = true
  }
}

// autoscaling enabled pool of application instances with version 2
resource "avi_pool" "waf_pool_v2" {
  name                = "waf_pool_v2"
  health_monitor_refs = [avi_healthmonitor.test_hm_1.id]
  tenant_ref          = data.avi_tenant.default_tenant.id
  cloud_ref           = data.avi_cloud.default_cloud.id
  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }
  ignore_servers = true
  analytics_policy {
    enable_realtime_metrics = true
  }
}

// Version 1 of the application (in-service)
resource "avi_server" "test_server_p11" {
  ip       = var.avi_app_server_v1
  port     = "80"
  pool_ref = avi_pool.waf_pool_v1.id
  hostname = "waf_app_v1_s1"
}

// Version 2 of the application (canary)
resource "avi_server" "test_server_p21" {
  ip       = var.avi_app_server_v2
  port     = "80"
  pool_ref = avi_pool.waf_pool_v2.id
  hostname = "waf_app_v2_s1"
}

