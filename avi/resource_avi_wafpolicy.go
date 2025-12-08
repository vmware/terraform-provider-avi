// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceWafPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"allow_mode_delegation": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"allowlist": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceWafPolicyAllowlistSchema(),
		},
		"application_signatures": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceWafApplicationSignaturesSchema(),
		},
		"auto_update_crs": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"bypass_static_extensions": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"crs_overrides": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupOverridesSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable_streaming": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"failure_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_FAILURE_MODE_OPEN",
		},
		"fixed_sampling_rate": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"geo_db_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_MODE_DETECTION_ONLY",
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"paranoia_level": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_PARANOIA_LEVEL_LOW",
		},
		"positive_security_model": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceWafPositiveSecurityModelSchema(),
		},
		"post_crs_groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
		},
		"pre_crs_groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
		},
		"required_data_files": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafPolicyRequiredDataFileSchema(),
		},
		"sampling_mode": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_SAMPLING_MODE_NO_SAMPLING",
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"use_evaluation_mode_on_crs_update": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"waf_crs_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"waf_profile_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourceAviWafPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafPolicyCreate,
		Read:   ResourceAviWafPolicyRead,
		Update: resourceAviWafPolicyUpdate,
		Delete: resourceAviWafPolicyDelete,
		Schema: ResourceWafPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafPolicyImporter,
		},
	}
}

func ResourceWafPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafPolicySchema()
	return ResourceImporter(d, m, "wafpolicy", s)
}

func ResourceAviWafPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicySchema()
	err := APIRead(d, meta, "wafpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicySchema()
	err := APICreate(d, meta, "wafpolicy", s)
	if err == nil {
		err = ResourceAviWafPolicyRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicySchema()
	var err error
	err = APIUpdate(d, meta, "wafpolicy", s)
	if err == nil {
		err = ResourceAviWafPolicyRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "wafpolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
