// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceVirtualServiceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"active_standby_se_tag": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "ACTIVE_STANDBY_SE_1",
		},
		"advertise_down_vs": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"allow_invalid_client_cert": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"analytics_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceAnalyticsPolicySchema(),
		},
		"analytics_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"application_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"azure_availability_set": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"bgp_local_preference": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"bgp_num_as_path_prepend": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"bgp_peer_labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"bot_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"bulk_sync_kvcache": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"close_client_conn_on_config_update": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"cloud_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "CLOUD_NONE",
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"connections_rate_limit": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRateProfileSchema(),
		},
		"content_rewrite": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceContentRewriteProfileSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"csrf_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"delay_fairness": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dns_info": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsInfoSchema(),
		},
		"dns_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsPoliciesSchema(),
		},
		"east_west_placement": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_autogw": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"enable_rhi": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"enable_rhi_snat": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"enable_session": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"error_page_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"flow_dist": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "LOAD_AWARE",
		},
		"flow_label_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "NO_LABEL",
		},
		"fqdn": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"host_name_xlate": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"http_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceHTTPPoliciesSchema(),
		},
		"icap_request_profile_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ign_pool_net_reach": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"jwt_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceJWTValidationVsConfigSchema(),
		},
		"l4_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceL4PoliciesSchema(),
		},
		"ldap_vs_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceLDAPVSConfigSchema(),
		},
		"limit_doser": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"max_cps_per_client": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"microservice_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"min_pools_up": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"network_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"network_security_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"nsx_securitygroup": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"oauth_vs_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceOAuthVSConfigSchema(),
		},
		"performance_limits": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePerformanceLimitsSchema(),
		},
		"pool_group_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"pool_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"remove_listening_port_on_vs_down": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"requests_rate_limit": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRateProfileSchema(),
		},
		"revoke_vip_route": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"saml_sp_config": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSAMLSPConfigSchema(),
		},
		"scaleout_ecmp": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"se_group_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"server_network_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_metadata": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"service_pool_select": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServicePoolSelectorSchema(),
		},
		"services": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServiceSchema(),
		},
		"sideband_profile": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSidebandProfileSchema(),
		},
		"snat_ip": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"snat_ip6_addresses": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceIpAddrSchema(),
		},
		"sp_pool_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ssl_key_and_certificate_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"ssl_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"ssl_profile_selectors": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSSLProfileSelectorSchema(),
		},
		"ssl_sess_cache_avg_size": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1024",
			ValidateFunc: validateInteger,
		},
		"sso_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"static_dns_records": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsRecordSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"test_se_datastore_level_1_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"topology_policies": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsPoliciesSchema(),
		},
		"traffic_clone_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"traffic_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VS_TYPE_NORMAL",
		},
		"use_bridge_ip_as_vip": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"use_vip_as_snat": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vh_domain_name": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"vh_matches": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVHMatchSchema(),
		},
		"vh_parent_vs_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vh_type": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VS_TYPE_VH_SNI",
		},
		"vip": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVipSchema(),
		},
		"vrf_context_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vs_datascripts": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceVSDataScriptsSchema(),
		},
		"vsvip_cloud_config_cksum": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vsvip_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"waf_policy_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"weight": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
	}
}

func resourceAviVirtualService() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVirtualServiceCreate,
		Read:   ResourceAviVirtualServiceRead,
		Update: resourceAviVirtualServiceUpdate,
		Delete: resourceAviVirtualServiceDelete,
		Schema: ResourceVirtualServiceSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVirtualServiceImporter,
		},
	}
}

func ResourceVirtualServiceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVirtualServiceSchema()
	return ResourceImporter(d, m, "virtualservice", s)
}

func ResourceAviVirtualServiceRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVirtualServiceSchema()
	err := APIRead(d, meta, "virtualservice", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVirtualServiceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVirtualServiceSchema()
	err := APICreate(d, meta, "virtualservice", s)
	if err == nil {
		err = ResourceAviVirtualServiceRead(d, meta)
	}
	return err
}

func resourceAviVirtualServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVirtualServiceSchema()
	var err error
	err = APIUpdate(d, meta, "virtualservice", s)
	if err == nil {
		err = ResourceAviVirtualServiceRead(d, meta)
	}
	return err
}

func resourceAviVirtualServiceDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "virtualservice")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
