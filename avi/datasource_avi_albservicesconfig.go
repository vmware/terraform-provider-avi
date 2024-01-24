// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviALBServicesConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviALBServicesConfigRead,
		Schema: map[string]*schema.Schema{
			"app_signature_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAppSignatureConfigSchema(),
			},
			"asset_contact": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceALBServicesUserSchema(),
			},
			"case_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceCaseConfigSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"feature_opt_in_status": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePortalFeatureOptInSchema(),
			},
			"ip_reputation_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpReputationConfigSchema(),
			},
			"mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"polling_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"portal_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"saas_licensing_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSaasLicensingInfoSchema(),
			},
			"session_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePulseServicesSessionConfigSchema(),
			},
			"split_proxy_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProxyConfigurationSchema(),
			},
			"tenant_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePulseServicesTenantConfigSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_split_proxy": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_tls": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_agent_db_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceUserAgentDBConfigSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"waf_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWafCrsConfigSchema(),
			},
		},
	}
}
