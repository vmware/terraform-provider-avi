// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviWafPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafPolicyRead,
		Schema: map[string]*schema.Schema{
			"allow_mode_delegation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"allowlist": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWafPolicyAllowlistSchema(),
			},
			"application_signatures": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWafApplicationSignaturesSchema(),
			},
			"auto_update_crs": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"bypass_static_extensions": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crs_overrides": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafRuleGroupOverridesSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_streaming": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"failure_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fixed_sampling_rate": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"geo_db_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
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
			"paranoia_level": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"positive_security_model": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWafPositiveSecurityModelSchema(),
			},
			"post_crs_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"pre_crs_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"required_data_files": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafPolicyRequiredDataFileSchema(),
			},
			"sampling_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_evaluation_mode_on_crs_update": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"waf_crs_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"waf_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
