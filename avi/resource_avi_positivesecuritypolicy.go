// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourcePositiveSecurityPolicySchema() map[string]*schema.Schema {
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
		"enable_positive_security_rule_updates": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_regex_programming": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"positive_security_params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePositiveSecurityParamsSchema(),
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

func resourceAviPositiveSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPositiveSecurityPolicyCreate,
		Read:   ResourceAviPositiveSecurityPolicyRead,
		Update: resourceAviPositiveSecurityPolicyUpdate,
		Delete: resourceAviPositiveSecurityPolicyDelete,
		Schema: ResourcePositiveSecurityPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePositiveSecurityPolicyImporter,
		},
	}
}

func ResourcePositiveSecurityPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePositiveSecurityPolicySchema()
	return ResourceImporter(d, m, "positivesecuritypolicy", s)
}

func ResourceAviPositiveSecurityPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePositiveSecurityPolicySchema()
	err := APIRead(d, meta, "positivesecuritypolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPositiveSecurityPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePositiveSecurityPolicySchema()
	err := APICreate(d, meta, "positivesecuritypolicy", s)
	if err == nil {
		err = ResourceAviPositiveSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviPositiveSecurityPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePositiveSecurityPolicySchema()
	var err error
	err = APIUpdate(d, meta, "positivesecuritypolicy", s)
	if err == nil {
		err = ResourceAviPositiveSecurityPolicyRead(d, meta)
	}
	return err
}

func resourceAviPositiveSecurityPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "positivesecuritypolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
