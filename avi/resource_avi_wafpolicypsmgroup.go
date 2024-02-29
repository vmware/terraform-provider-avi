// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceWafPolicyPSMGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"enable": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"hit_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_ACTION_ALLOW_PARAMETER",
		},
		"is_learning_group": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"locations": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafPSMLocationSchema(),
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"miss_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "WAF_ACTION_NO_OP",
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
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
	}
}

func resourceAviWafPolicyPSMGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafPolicyPSMGroupCreate,
		Read:   ResourceAviWafPolicyPSMGroupRead,
		Update: resourceAviWafPolicyPSMGroupUpdate,
		Delete: resourceAviWafPolicyPSMGroupDelete,
		Schema: ResourceWafPolicyPSMGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafPolicyPSMGroupImporter,
		},
	}
}

func ResourceWafPolicyPSMGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafPolicyPSMGroupSchema()
	return ResourceImporter(d, m, "wafpolicypsmgroup", s)
}

func ResourceAviWafPolicyPSMGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicyPSMGroupSchema()
	err := APIRead(d, meta, "wafpolicypsmgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafPolicyPSMGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicyPSMGroupSchema()
	err := APICreate(d, meta, "wafpolicypsmgroup", s)
	if err == nil {
		err = ResourceAviWafPolicyPSMGroupRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyPSMGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafPolicyPSMGroupSchema()
	var err error
	err = APIUpdate(d, meta, "wafpolicypsmgroup", s)
	if err == nil {
		err = ResourceAviWafPolicyPSMGroupRead(d, meta)
	}
	return err
}

func resourceAviWafPolicyPSMGroupDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "wafpolicypsmgroup")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
