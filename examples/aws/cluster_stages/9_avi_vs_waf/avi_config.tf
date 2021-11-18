
data "avi_cloud" "tf_cloud" {
  name = var.vs_cloud
}

data "avi_healthmonitor" "system_http_hm" {
  name = "System-HTTP"
}

data "avi_applicationprofile" "system_http" {
  name = "System-HTTP"
}

data "avi_wafpolicy" "wafpol_tf_vs" {
  name = "custom-waf-policy"
}

data "aws_vpc" "avi_vpc" {
  id = var.aws_vpc_id
}
data "aws_subnet" "terraform-subnets-0" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-subnet-0"]
  }
}
data "aws_subnet" "terraform-subnets-1" {
  filter {
    name   = "tag:Name"
    values = ["${var.project_name}-terraform-subnet-1"]
  }
}

resource "avi_vsvip" "vsvip_tf_vs" {
  cloud_ref = data.avi_cloud.tf_cloud.id
  name      = "vsvip-${var.vs_name}-${var.vs_cloud}"
  dns_info {
    fqdn                    = var.vs_fqdn
    num_records_in_response = 0
    ttl                     = 15
  }

  vip {
    auto_allocate_floating_ip = true
    auto_allocate_ip          = true
    availability_zone         = var.aws_availability_zone
    enabled                   = true
    subnet_uuid               = var.subnet_id
    vip_id                    = "1"
  }
}

resource "avi_pool" "pool_tf_vs" {
  cloud_ref           = data.avi_cloud.tf_cloud.id
  lb_algorithm        = "LB_ALGORITHM_ROUND_ROBIN"
  default_server_port = 80
  health_monitor_refs = [
    data.avi_healthmonitor.system_http_hm.id,
  ]
  name = "${var.vs_name}-pool"

  fail_action {
    type = "FAIL_ACTION_CLOSE_CONN"
  }

  analytics_policy {
    enable_realtime_metrics = true
  }
  servers {
    ip {
      type = "V4"
      addr = "10.2.11.47"
    }
  }
}

resource "avi_virtualservice" "vs_tf_vs" {
  application_profile_ref = data.avi_applicationprofile.system_http.id
  cloud_ref               = data.avi_cloud.tf_cloud.id
  cloud_type              = "CLOUD_AWS"
  name                    = var.vs_name
  pool_ref                = avi_pool.pool_tf_vs.id
  vsvip_ref               = avi_vsvip.vsvip_tf_vs.id

  waf_policy_ref = data.avi_wafpolicy.wafpol_tf_vs.id

  analytics_policy {
    all_headers              = true
    significant_log_throttle = 0
    udf_log_throttle         = 0

    full_client_logs {
      duration = 0
      enabled  = true
      throttle = 0
    }

    metrics_realtime_update {
      duration = 0
      enabled  = true
    }
  }

  services {
    enable_http2 = false
    enable_ssl   = false
    port         = 80
  }
}
