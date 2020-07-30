/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviWafPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafPolicyRead,
		Schema: map[string]*schema.Schema{
			"allow_mode_delegation": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"application_signatures": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWafApplicationSignaturesSchema(),
			},
			"confidence_override": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAppLearningConfidenceOverrideSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crs_groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_app_learning": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_auto_rule_updates": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_regex_learning": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"failure_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"learning_params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAppLearningParamsSchema(),
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
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
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
			"whitelist": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWafPolicyWhitelistSchema(),
			},
		},
	}
}
