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
				Type:     schema.TypeString,
				Computed: true,
			},
			"capacity_estimation_ttfb_thresh": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_server_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delete_server_on_dns_refresh": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_http2": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"graceful_disable_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gslb_sp_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"health_monitor_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"horizon_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHorizonProfileSchema(),
			},
			"host_check_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"http2_properties": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTP2PoolPropertiesSchema(),
			},
			"ignore_server_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"inline_health_monitor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipaddrgroup_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lb_algo_rr_per_se": {
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"lb_algorithm_hash": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lookup_server_by_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"max_concurrent_connections_per_server": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_conn_rate_per_server": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"min_health_monitors_up": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_servers_up": {
				Type:     schema.TypeString,
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
			"pool_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"request_queue_depth": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"request_queue_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resolve_pool_by_dns": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rewrite_host_header_to_server_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rewrite_host_header_to_sni": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"routing_pool": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_service_ssl_mode": {
				Type:     schema.TypeString,
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
