// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviPool() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPoolRead,
		Schema: map[string]*schema.Schema{
			"ignore_servers": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"analytics_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePoolAnalyticsPolicySchema(),
			},
			"analytics_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"append_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_persistence_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"autoscale_launch_config_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"autoscale_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"autoscale_policy_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"capacity_estimation": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"capacity_estimation_ttfb_thresh": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"cloud_config_cksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"conn_pool_properties": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConnPoolPropertiesSchema(),
			},
			"connection_ramp_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_server_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"delete_server_on_dns_refresh": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_name": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"east_west": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_http2": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"external_autoscale_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"fail_action": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceFailActionSchema(),
			},
			"fewest_tasks_feedback_delay": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"graceful_disable_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"health_monitor_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"host_check_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"http2_properties": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTP2PoolPropertiesSchema(),
			},
			"ignore_server_port": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"inline_health_monitor": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ipaddrgroup_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lb_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lb_algorithm_consistent_hash_hdr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lb_algorithm_core_nonaffinity": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"lb_algorithm_hash": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lookup_server_by_name": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"max_concurrent_connections_per_server": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_conn_rate_per_server": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"min_health_monitors_up": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_servers_up": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceNetworkFilterSchema(),
			},
			"nsx_securitygroup": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pki_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"placement_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourcePlacementNetworkSchema(),
			},
			"request_queue_depth": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"request_queue_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"resolve_pool_by_dns": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"rewrite_host_header_to_server_name": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"rewrite_host_header_to_sni": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"routing_pool": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"server_disable_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_reselect": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTPServerReselectSchema(),
			},
			"server_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceServerSchema(),
			},
			"service_metadata": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sni_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ssl_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier1_lr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_service_port": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"use_service_ssl_mode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vrf_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
