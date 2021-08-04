// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceAnalyticsProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apdex_response_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "500",
			ValidateFunc: validateInteger,
		},
		"apdex_response_tolerated_factor": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateFloat,
		},
		"apdex_rtt_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "250",
			ValidateFunc: validateInteger,
		},
		"apdex_rtt_tolerated_factor": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateFloat,
		},
		"apdex_rum_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5000",
			ValidateFunc: validateInteger,
		},
		"apdex_rum_tolerated_factor": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateFloat,
		},
		"apdex_server_response_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "400",
			ValidateFunc: validateInteger,
		},
		"apdex_server_response_tolerated_factor": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateFloat,
		},
		"apdex_server_rtt_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "125",
			ValidateFunc: validateInteger,
		},
		"apdex_server_rtt_tolerated_factor": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateFloat,
		},
		"client_log_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceClientLogConfigurationSchema(),
		},
		"client_log_streaming_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceClientLogStreamingConfigSchema(),
		},
		"conn_lossy_ooo_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "50",
			ValidateFunc: validateInteger,
		},
		"conn_lossy_timeo_rexmt_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20",
			ValidateFunc: validateInteger,
		},
		"conn_lossy_total_rexmt_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "50",
			ValidateFunc: validateInteger,
		},
		"conn_lossy_zero_win_size_event_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2",
			ValidateFunc: validateInteger,
		},
		"conn_server_lossy_ooo_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "50",
			ValidateFunc: validateInteger,
		},
		"conn_server_lossy_timeo_rexmt_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20",
			ValidateFunc: validateInteger,
		},
		"conn_server_lossy_total_rexmt_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "50",
			ValidateFunc: validateInteger,
		},
		"conn_server_lossy_zero_win_size_event_threshold": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2",
			ValidateFunc: validateInteger,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable_adaptive_config": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"enable_advanced_analytics": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"enable_ondemand_metrics": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"enable_se_analytics": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"enable_server_analytics": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"enable_vs_analytics": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"exclude_client_close_before_request_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_dns_policy_drop_as_significant": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_gs_down_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_http_error_codes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"exclude_invalid_dns_domain_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_invalid_dns_query_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_issuer_revoked_ocsp_responses_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"exclude_no_dns_record_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_no_valid_gs_member_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_persistence_change_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_revoked_ocsp_responses_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"exclude_server_dns_error_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_server_tcp_reset_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_sip_error_codes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"exclude_stale_ocsp_responses_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"exclude_syn_retransmit_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_tcp_reset_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"exclude_unavailable_ocsp_responses_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"exclude_unsupported_dns_query_as_error": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"healthscore_max_server_limit": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20",
			ValidateFunc: validateInteger,
		},
		"hs_event_throttle_window": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1209600",
			ValidateFunc: validateInteger,
		},
		"hs_max_anomaly_penalty": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateInteger,
		},
		"hs_max_resources_penalty": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "25",
			ValidateFunc: validateInteger,
		},
		"hs_max_security_penalty": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "100",
			ValidateFunc: validateInteger,
		},
		"hs_min_dos_rate": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1000",
			ValidateFunc: validateInteger,
		},
		"hs_performance_boost": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"hs_pscore_traffic_threshold_l4_client": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateFloat,
		},
		"hs_pscore_traffic_threshold_l4_server": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "10",
			ValidateFunc: validateFloat,
		},
		"hs_security_certscore_expired": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateFloat,
		},
		"hs_security_certscore_gt30d": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateFloat,
		},
		"hs_security_certscore_le07d": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2",
			ValidateFunc: validateFloat,
		},
		"hs_security_certscore_le30d": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateFloat,
		},
		"hs_security_chain_invalidity_penalty": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateFloat,
		},
		"hs_security_cipherscore_eq000b": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateFloat,
		},
		"hs_security_cipherscore_ge128b": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateFloat,
		},
		"hs_security_cipherscore_lt128b": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3.5",
			ValidateFunc: validateFloat,
		},
		"hs_security_encalgo_score_none": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateFloat,
		},
		"hs_security_encalgo_score_rc4": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "2.5",
			ValidateFunc: validateFloat,
		},
		"hs_security_hsts_penalty": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateFloat,
		},
		"hs_security_nonpfs_penalty": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateFloat,
		},
		"hs_security_ocsp_revoked_score": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateFloat,
		},
		"hs_security_selfsignedcert_penalty": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateFloat,
		},
		"hs_security_ssl30_score": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3.5",
			ValidateFunc: validateFloat,
		},
		"hs_security_tls10_score": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateFloat,
		},
		"hs_security_tls11_score": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateFloat,
		},
		"hs_security_tls12_score": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateFloat,
		},
		"hs_security_weak_signature_algo_penalty": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateFloat,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ondemand_metrics_idle_timeout": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1800",
			ValidateFunc: validateInteger,
		},
		"ranges": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceHTTPStatusRangeSchema(),
		},
		"resp_code_block": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"sensitive_log_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSensitiveLogProfileSchema(),
		},
		"sip_log_depth": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "20",
			ValidateFunc: validateInteger,
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
	}
}

func resourceAviAnalyticsProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAnalyticsProfileCreate,
		Read:   ResourceAviAnalyticsProfileRead,
		Update: resourceAviAnalyticsProfileUpdate,
		Delete: resourceAviAnalyticsProfileDelete,
		Schema: ResourceAnalyticsProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAnalyticsProfileImporter,
		},
	}
}

func ResourceAnalyticsProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAnalyticsProfileSchema()
	return ResourceImporter(d, m, "analyticsprofile", s)
}

func ResourceAviAnalyticsProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAnalyticsProfileSchema()
	err := APIRead(d, meta, "analyticsprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAnalyticsProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAnalyticsProfileSchema()
	err := APICreateOrUpdate(d, meta, "analyticsprofile", s)
	if err == nil {
		err = ResourceAviAnalyticsProfileRead(d, meta)
	}
	return err
}

func resourceAviAnalyticsProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAnalyticsProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "analyticsprofile", s)
	if err == nil {
		err = ResourceAviAnalyticsProfileRead(d, meta)
	}
	return err
}

func resourceAviAnalyticsProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "analyticsprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviAnalyticsProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
