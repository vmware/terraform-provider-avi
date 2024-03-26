// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceCSRFPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"cookie_name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "X-CSRF-TOKEN",
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"rules": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceCSRFRuleSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"token_validity_time_min": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "360",
			ValidateFunc: validateInteger,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviCSRFPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCSRFPolicyCreate,
		Read:   ResourceAviCSRFPolicyRead,
		Update: resourceAviCSRFPolicyUpdate,
		Delete: resourceAviCSRFPolicyDelete,
		Schema: ResourceCSRFPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCSRFPolicyImporter,
		},
	}
}

func ResourceCSRFPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCSRFPolicySchema()
	return ResourceImporter(d, m, "csrfpolicy", s)
}

func ResourceAviCSRFPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCSRFPolicySchema()
	err := APIRead(d, meta, "csrfpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCSRFPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCSRFPolicySchema()
	err := APICreate(d, meta, "csrfpolicy", s)
	if err == nil {
		err = ResourceAviCSRFPolicyRead(d, meta)
	}
	return err
}

func resourceAviCSRFPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCSRFPolicySchema()
	var err error
	err = APIUpdate(d, meta, "csrfpolicy", s)
	if err == nil {
		err = ResourceAviCSRFPolicyRead(d, meta)
	}
	return err
}

func resourceAviCSRFPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "csrfpolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
