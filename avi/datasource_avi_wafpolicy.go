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
			"confidence_override": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAppLearningConfidenceOverrideSchema(),
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
			"enable_app_learning": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_auto_rule_updates": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_regex_learning": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"failure_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"geo_db_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"learning_params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAppLearningParamsSchema(),
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"min_confidence": {
				Type:     schema.TypeString,
				Computed: true,
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
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"updated_crs_rules_in_detection_mode": {
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
