// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviVirtualService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVirtualServiceRead,
		Schema: map[string]*schema.Schema{
			"active_standby_se_tag": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"advertise_down_vs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allow_invalid_client_cert": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"analytics_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAnalyticsPolicySchema(),
			},
			"analytics_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bgp_peer_labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"bot_policy_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bulk_sync_kvcache": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"close_client_conn_on_config_update": {
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
			"cloud_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"connections_rate_limit": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"content_rewrite": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceContentRewriteProfileSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"delay_fairness": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsInfoSchema(),
			},
			"dns_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsPoliciesSchema(),
			},
			"east_west_placement": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_autogw": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_rhi": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_rhi_snat": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error_page_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_dist": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flow_label_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fqdn": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_name_xlate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceHTTPPoliciesSchema(),
			},
			"icap_request_profile_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ign_pool_net_reach": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"jwt_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceJWTValidationVsConfigSchema(),
			},
			"l4_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceL4PoliciesSchema(),
			},
			"ldap_vs_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceLDAPVSConfigSchema(),
			},
			"limit_doser": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"max_cps_per_client": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"microservice_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_pools_up": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"network_security_policy_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nsx_securitygroup": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"performance_limits": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePerformanceLimitsSchema(),
			},
			"pool_group_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pool_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remove_listening_port_on_vs_down": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"requests_rate_limit": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRateProfileSchema(),
			},
			"saml_sp_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSAMLSPConfigSchema(),
			},
			"scaleout_ecmp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_group_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_policy_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_network_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_metadata": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_pool_select": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceServicePoolSelectorSchema(),
			},
			"services": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceServiceSchema(),
			},
			"sideband_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSidebandProfileSchema(),
			},
			"snat_ip": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"ssl_key_and_certificate_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ssl_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_profile_selectors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceSSLProfileSelectorSchema(),
			},
			"ssl_sess_cache_avg_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sso_policy_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static_dns_records": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsRecordSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"test_se_datastore_level_1_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"topology_policies": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsPoliciesSchema(),
			},
			"traffic_clone_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"traffic_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_bridge_ip_as_vip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_vip_as_snat": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vh_domain_name": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vh_matches": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceVHMatchSchema(),
			},
			"vh_parent_vs_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vh_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vip": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceVipSchema(),
			},
			"vrf_context_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_datascripts": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceVSDataScriptsSchema(),
			},
			"vsvip_cloud_config_cksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsvip_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"waf_policy_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"weight": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
