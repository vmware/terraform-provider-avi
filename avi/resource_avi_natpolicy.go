// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceNatPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceNatRuleSchema(),
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

func resourceAviNatPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNatPolicyCreate,
		Read:   ResourceAviNatPolicyRead,
		Update: resourceAviNatPolicyUpdate,
		Delete: resourceAviNatPolicyDelete,
		Schema: ResourceNatPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNatPolicyImporter,
		},
	}
}

func ResourceNatPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNatPolicySchema()
	return ResourceImporter(d, m, "natpolicy", s)
}

func ResourceAviNatPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNatPolicySchema()
	err := APIRead(d, meta, "natpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNatPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNatPolicySchema()
	err := APICreate(d, meta, "natpolicy", s)
	if err == nil {
		err = ResourceAviNatPolicyRead(d, meta)
	}
	return err
}

func resourceAviNatPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNatPolicySchema()
	var err error
	err = APIUpdate(d, meta, "natpolicy", s)
	if err == nil {
		err = ResourceAviNatPolicyRead(d, meta)
	}
	return err
}

func resourceAviNatPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "natpolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
