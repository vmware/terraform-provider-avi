/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceAnalyticsProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"apdex_response_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  500,
		},
		"apdex_response_tolerated_factor": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "4.0",
		},
		"apdex_rtt_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  250,
		},
		"apdex_rtt_tolerated_factor": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "4.0",
		},
		"apdex_rum_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  5000,
		},
		"apdex_rum_tolerated_factor": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "4.0",
		},
		"apdex_server_response_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  400,
		},
		"apdex_server_response_tolerated_factor": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "4.0",
		},
		"apdex_server_rtt_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  125,
		},
		"apdex_server_rtt_tolerated_factor": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "4.0",
		},
		"client_log_config": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceClientLogConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"client_log_streaming_config": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceClientLogStreamingConfigSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"conn_lossy_ooo_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  50,
		},
		"conn_lossy_timeo_rexmt_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"conn_lossy_total_rexmt_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  50,
		},
		"conn_lossy_zero_win_size_event_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"conn_server_lossy_ooo_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  50,
		},
		"conn_server_lossy_timeo_rexmt_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"conn_server_lossy_total_rexmt_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  50,
		},
		"conn_server_lossy_zero_win_size_event_threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  2,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"disable_ondemand_metrics": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disable_se_analytics": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disable_server_analytics": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"disable_vs_analytics": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"enable_advanced_analytics": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"exclude_client_close_before_request_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_dns_policy_drop_as_significant": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_gs_down_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_http_error_codes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"exclude_invalid_dns_domain_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_invalid_dns_query_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_no_dns_record_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_no_valid_gs_member_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_persistence_change_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_server_dns_error_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_server_tcp_reset_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_sip_error_codes": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeInt},
		},
		"exclude_syn_retransmit_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_tcp_reset_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"exclude_unsupported_dns_query_as_error": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"healthscore_max_server_limit": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"hs_event_throttle_window": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1209600,
		},
		"hs_max_anomaly_penalty": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"hs_max_resources_penalty": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  25,
		},
		"hs_max_security_penalty": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  100,
		},
		"hs_min_dos_rate": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1000,
		},
		"hs_performance_boost": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  0,
		},
		"hs_pscore_traffic_threshold_l4_client": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "10.0",
		},
		"hs_pscore_traffic_threshold_l4_server": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "10.0",
		},
		"hs_security_certscore_expired": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "0.0",
		},
		"hs_security_certscore_gt30d": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "5.0",
		},
		"hs_security_certscore_le07d": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "2.0",
		},
		"hs_security_certscore_le30d": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "4.0",
		},
		"hs_security_chain_invalidity_penalty": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "1.0",
		},
		"hs_security_cipherscore_eq000b": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "0.0",
		},
		"hs_security_cipherscore_ge128b": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "5.0",
		},
		"hs_security_cipherscore_lt128b": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "3.5",
		},
		"hs_security_encalgo_score_none": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "0.0",
		},
		"hs_security_encalgo_score_rc4": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "2.5",
		},
		"hs_security_hsts_penalty": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "1.0",
		},
		"hs_security_nonpfs_penalty": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "1.0",
		},
		"hs_security_selfsignedcert_penalty": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "1.0",
		},
		"hs_security_ssl30_score": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "3.5",
		},
		"hs_security_tls10_score": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "5.0",
		},
		"hs_security_tls11_score": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "5.0",
		},
		"hs_security_tls12_score": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "5.0",
		},
		"hs_security_weak_signature_algo_penalty": &schema.Schema{
			Type:     schema.TypeFloat,
			Optional: true,
			Default:  "1.0",
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"ondemand_metrics_idle_timeout": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1800,
		},
		"ranges": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceHTTPStatusRangeSchema(),
		},
		"resp_code_block": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"sensitive_log_profile": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSensitiveLogProfileSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"sip_log_depth": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  20,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": &schema.Schema{
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
	err := ApiRead(d, meta, "analyticsprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAnalyticsProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAnalyticsProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "analyticsprofile", s)
	if err == nil {
		err = ResourceAviAnalyticsProfileRead(d, meta)
	}
	return err
}

func resourceAviAnalyticsProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAnalyticsProfileSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "analyticsprofile", s)
	if err == nil {
		err = ResourceAviAnalyticsProfileRead(d, meta)
	}
	return err
}

func resourceAviAnalyticsProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "analyticsprofile"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
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
