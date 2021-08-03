// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviAnalyticsProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAnalyticsProfileRead,
		Schema: map[string]*schema.Schema{
			"apdex_response_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apdex_response_tolerated_factor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apdex_rtt_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apdex_rtt_tolerated_factor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apdex_rum_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apdex_rum_tolerated_factor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apdex_server_response_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apdex_server_response_tolerated_factor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apdex_server_rtt_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"apdex_server_rtt_tolerated_factor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_log_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceClientLogConfigurationSchema(),
			},
			"client_log_streaming_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceClientLogStreamingConfigSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"conn_lossy_ooo_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"conn_lossy_timeo_rexmt_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"conn_lossy_total_rexmt_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"conn_lossy_zero_win_size_event_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"conn_server_lossy_ooo_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"conn_server_lossy_timeo_rexmt_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"conn_server_lossy_total_rexmt_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"conn_server_lossy_zero_win_size_event_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_adaptive_config": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_advanced_analytics": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_ondemand_metrics": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_se_analytics": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_server_analytics": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_vs_analytics": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_client_close_before_request_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_dns_policy_drop_as_significant": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_gs_down_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_http_error_codes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"exclude_invalid_dns_domain_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_invalid_dns_query_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_issuer_revoked_ocsp_responses_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_no_dns_record_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_no_valid_gs_member_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_persistence_change_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_revoked_ocsp_responses_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_server_dns_error_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_server_tcp_reset_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_sip_error_codes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeInt},
			},
			"exclude_stale_ocsp_responses_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_syn_retransmit_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_tcp_reset_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_unavailable_ocsp_responses_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"exclude_unsupported_dns_query_as_error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"healthscore_max_server_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_event_throttle_window": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_max_anomaly_penalty": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_max_resources_penalty": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_max_security_penalty": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_min_dos_rate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_performance_boost": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_pscore_traffic_threshold_l4_client": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_pscore_traffic_threshold_l4_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_certscore_expired": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_certscore_gt30d": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_certscore_le07d": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_certscore_le30d": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_chain_invalidity_penalty": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_cipherscore_eq000b": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_cipherscore_ge128b": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_cipherscore_lt128b": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_encalgo_score_none": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_encalgo_score_rc4": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_hsts_penalty": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_nonpfs_penalty": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_ocsp_revoked_score": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_selfsignedcert_penalty": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_ssl30_score": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_tls10_score": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_tls11_score": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_tls12_score": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hs_security_weak_signature_algo_penalty": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"latency_audit_props": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceLatencyAuditPropertiesSchema(),
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ondemand_metrics_idle_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceHTTPStatusRangeSchema(),
			},
			"resp_code_block": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"sensitive_log_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSensitiveLogProfileSchema(),
			},
			"sip_log_depth": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
