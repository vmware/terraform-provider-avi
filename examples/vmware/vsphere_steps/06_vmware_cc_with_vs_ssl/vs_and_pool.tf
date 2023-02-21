resource "avi_pool" "pool_tf_vs" {
  cloud_ref           = avi_cloud.vmware_cloud_tf.id
  default_server_port = var.default_port
  pool_type    = "POOL_TYPE_GENERIC_APP"
  #health_monitor_refs = [
  #  avi_healthmonitor.cop_healthmonitor.id,
  #]
  name = var.pool_name

  rewrite_host_header_to_sni = var.rewrite_host_header
  server_name                = var.server_name

  servers {
    hostname              = var.hostname
    ratio                 = 1
    enabled               = true
    resolve_server_by_dns = var.resolve_server_by_dns
    ip {
      addr = var.server_ip
      type = "V4"
    }
  }
    fail_action {
    type = var.fail_type
  }

  analytics_policy {
    enable_realtime_metrics = var.realtime_metrics
  }

  # Specifying an SSL/TLS profile reference for the pool enables TLS to the back-end
  # mTLS will be used to authenticate the Avi SE to the pool members
  #ssl_profile_ref             = avi_sslprofile.cop_sslprofile.id
  #ssl_key_and_certificate_ref = data.avi_sslkeyandcertificate.cop_clientcertificate.id

  #lifecycle {
  #  ignore_changes = [
  #    servers,
  #  ]
  #}
}

resource "avi_virtualservice" "https_vs" {
  name                           =  "vsdemo"
  pool_ref                       = avi_pool.pool_tf_vs.id
  tenant_ref                     = "admin"
  vsvip_ref                      = avi_vsvip.avi_vsvip.id
  cloud_ref                      = avi_cloud.vmware_cloud_tf.id
  application_profile_ref        = avi_applicationprofile.application_profile.id
  network_profile_ref            = avi_networkprofile.network_profile.id
  services {
    port           = var.vs_port
    enable_ssl     = true
  }
}
resource "avi_applicationprofile" "application_profile" {
  name = "terraform_https_application_profile"
  type = "APPLICATION_PROFILE_TYPE_HTTP"
  http_profile {
      hsts_enabled = true
      secure_cookie_enabled = true
      httponly_enabled = true
      http_to_https = true
      server_side_redirect_to_https = true
      x_forwarded_proto_enabled = true
      compression_profile {
        compression = false
        remove_accept_encoding_header = true
        type = "AUTO_COMPRESSION"
        compressible_content_ref = avi_stringgroup.string_group.id
      }
      post_accept_timeout = 30000
      client_header_timeout = 10000
      client_body_timeout = 30000
      keepalive_timeout = 30000
      client_max_header_size = 12
      client_max_request_size = 48
      client_max_body_size = 0
      cache_config {
        enabled = false
        xcache_header = true
        age_header = true
        date_header = true
        min_object_size = 100
        max_object_size = 4194304
        default_expire = 600
        heuristic_expire = false
        max_cache_size = 0
        query_cacheable = false
        aggressive = false
        ignore_request_cache_control = false
    }
  }
  tenant_ref   = var.tenant
}

resource "avi_stringgroup" "string_group" {
  name = "System-Compressible-Content-Types"
  kv {
      key = "text/html"
  }
  kv {
      key = "text/xml"
  }
  type = "SG_TYPE_STRING"
  longest_match = false
  tenant_ref   = var.tenant
}

resource "avi_sslprofile" "ssl_profile" {
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
resource "avi_networkprofile" "network_profile" {
  name = "tf-network-profile"
  profile {
    type = "PROTOCOL_TYPE_TCP_PROXY"
    tcp_proxy_profile {
      ignore_time_wait = false
      time_wait_delay = 2000
      max_retransmissions = 8
      max_syn_retransmissions = 8
      automatic = true
      receive_window = 64
      nagles_algorithm = false
      ip_dscp = 0
      reorder_threshold = 10
      min_rexmt_timeout = 100
      idle_connection_type = "KEEP_ALIVE"
      idle_connection_timeout = 600
      use_interface_mtu = true
      cc_algo = "CC_ALGO_NEW_RENO"
      aggressive_congestion_avoidance = false
      slow_start_scaling_factor = 1
      congestion_recovery_scaling_factor = 2
      reassembly_queue_size = 0
      keepalive_in_halfclose_state = true
      auto_window_growth = true
    }
  }
  connection_mirror = false
  tenant_ref = var.tenant
}
