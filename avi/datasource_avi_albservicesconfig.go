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
			"polling_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"portal_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proactive_support_defaults": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProactiveSupportDefaultsSchema(),
			},
			"split_proxy_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProxyConfigurationSchema(),
			},
			"use_split_proxy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"use_tls": {
				Type:     schema.TypeBool,
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
