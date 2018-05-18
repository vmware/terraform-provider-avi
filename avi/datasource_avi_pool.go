/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPool() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPoolRead,
		Schema: map[string]*schema.Schema{
			"a_pool": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ab_pool": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAbPoolSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ab_priority": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"apic_epg_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"application_persistence_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"autoscale_launch_config_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"autoscale_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"autoscale_policy_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"capacity_estimation": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"capacity_estimation_ttfb_thresh": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"cloud_config_cksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"connection_ramp_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"default_server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  80,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_name": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"east_west": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"external_autoscale_groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fail_action": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceFailActionSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"fewest_tasks_feedback_delay": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"graceful_disable_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"health_monitor_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"host_check_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"inline_health_monitor": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ipaddrgroup_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lb_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LB_ALGORITHM_LEAST_CONNECTIONS"},
			"lb_algorithm_consistent_hash_hdr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"lb_algorithm_core_nonaffinity": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"lb_algorithm_hash": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "LB_ALGORITHM_CONSISTENT_HASH_SOURCE_IP_ADDRESS"},
			"lookup_server_by_name": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"max_concurrent_connections_per_server": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"max_conn_rate_per_server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceRateProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceNetworkFilterSchema(),
			},
			"nsx_securitygroup": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pki_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"placement_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePlacementNetworkSchema(),
			},
			"prst_hdr_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"request_queue_depth": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  128,
			},
			"request_queue_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewrite_host_header_to_server_name": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewrite_host_header_to_sni": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_auto_scale": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"server_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_reselect": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPServerReselectSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceServerSchema(),
			},
			"sni_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ssl_key_and_certificate_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ssl_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_service_port": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"vrf_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
