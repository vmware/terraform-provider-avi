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

data "avi_wafapplicationsignatureprovider" "signature_provider" {
    name = "System-WafApplicationSignatures-Trustwave"
}

data "avi_sslkeyandcertificate" "ssl_cert1" {
	name = "System-Default-Cert"
}

data "avi_applicationprofile" "application_profile1" {
	name = "System-HTTP"
}


data "avi_wafcrs" "waf_crs" {
    name = var.waf_crs_name
}

data "avi_wafprofile" "waf_profile" {
  name = var.waf_profile
}

resource "avi_cloud" "cloud_none" {
  vtype                        = "CLOUD_NONE"
  license_tier                 = "ENTERPRISE"
  dhcp_enabled                 = false
  prefer_static_routes         = false
  license_type                 = "LIC_CORES"
  mtu                          = 1500
  state_based_dns_registration = true
  name                         = "test-tf-cloud"
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_networkprofile" "network_profile1" {
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
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_stringgroup" "stringgroup1" {
  name = "System-Compressible-Content-Types"
  kv {
      key = "text/html"
  }
  kv {
      key = "text/xml"
  }
  type = "SG_TYPE_STRING"
  longest_match = false
  tenant_ref = data.avi_tenant.default_tenant.id
}

resource "avi_sslprofile" "ssl_profile1" {
  name = "tf-ssl-profile"
  accepted_ciphers = "ECDHE-ECDSA-AES128-GCM-SHA256:ECDHE-ECDSA-AES256-GCM-SHA384:ECDHE-RSA-AES128-GCM-SHA256:ECDHE-RSA-AES256-GCM-SHA384"
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
  tenant_ref = data.avi_tenant.default_tenant.id
  cloud_ref = avi_cloud.cloud_none.id
}

resource "avi_wafpolicy" "waf_app_policy" {
    name = "test_waf_app_policy"
    tenant_ref = data.avi_tenant.default_tenant.id
    waf_profile_ref = data.avi_wafprofile.waf_profile.id
    waf_crs_ref = data.avi_wafcrs.waf_crs.id
    mode = "WAF_MODE_DETECTION_ONLY"
    paranoia_level = "WAF_PARANOIA_LEVEL_LOW"
    failure_mode = "WAF_FAILURE_MODE_OPEN"
    allow_mode_delegation = true
    enable_app_learning = false
    application_signatures {
        provider_ref = data.avi_wafapplicationsignatureprovider.signature_provider.id
    }
    learning_params {
        sampling_percent = 1
        update_interval = 30
        max_uris = 500
        max_params = 100
        enable_per_uri_learning = true
        min_hits_to_learn = 10000
    }
    confidence_override {
        confid_very_high_value = 9999
        confid_high_value = 9500
        confid_probable_value = 9000
        confid_low_value = 7500
    }
    enable_auto_rule_updates = true
    enable_regex_learning = false
    # Override attributes for CRS rules
    crs_overrides {
        # The name of the group where attributes or rules are overridden
        name = "CRS_402_Additional_Rules"
        # Override the enable flag for this group
        enable = true
        # Replace the exclude list for this group
        exclude_list {
            # Client IP Subnet to exclude for WAF rules
            client_subnet {
                ip_addr {
                    addr = "10.20.30.40"
                    type = "V4"
                }
                mask = 24
            }
            # URI Path to exclude for WAF rules
            uri_path = "abc.com"
            # These match_elements in the HTTP Transaction (if present) will be excluded when executing WAF Rules
            match_element = "REQUEST_URI"
            # Criteria for URI matching
            uri_match_criteria {
                # String Operation to use for matching the Exclusion
                match_op = "EQUALS"
                # Case sensitivity to use for the matching
                match_case = "SENSITIVE"
            }
            # Criteria for match_element matching
            match_element_criteria {
              match_op = "EQUALS"
              match_case = "SENSITIVE"
            }
            description = "Free-text comment about this exclusion"

        }
        # Override attributes of application signature rules
         rule_overrides { 
            # Override the enable flag for this rule
            enable = true
            # The rule_id of the rule where attributes are overridden
            rule_id = "4022020"
            # Override the waf mode for this rule
            mode = "WAF_MODE_ENFORCEMENT"
            exclude_list {
                client_subnet {
                    ip_addr {
                        addr = "10.20.21.22"
                        type = "V4"
                    }
                    mask = 24
                }
                uri_path = "xyz.com"
                match_element = "REQUEST_BODY"
                uri_match_criteria {
                    match_op = "EQUALS"
                    match_case = "SENSITIVE"
                }
                match_element_criteria {
                    match_op = "EQUALS"
                    match_case = "SENSITIVE"
                }
            }
        }
    }
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


resource "avi_virtualservice" "https_vs" {
  name                          = var.vs_name
  pool_group_ref                = avi_poolgroup.poolgroup1.id
  tenant_ref                    = data.avi_tenant.default_tenant.id
  vsvip_ref                     = avi_vsvip.test_vsvip.id
  cloud_ref                     = avi_cloud.cloud_none.id
  ssl_key_and_certificate_refs  = [data.avi_sslkeyandcertificate.ssl_cert1.id]
  ssl_profile_ref               = avi_sslprofile.ssl_profile1.id
  application_profile_ref       = data.avi_applicationprofile.application_profile1.id
  network_profile_ref           = avi_networkprofile.network_profile1.id
  waf_policy_ref				= avi_wafpolicy.waf_app_policy.id
  services {
    port           = var.vs_port
    enable_ssl     = true
  }
}
