// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourcePasswordPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"lockout_evaluation_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "900",
			ValidateFunc: validateInteger,
		},
		"lockout_max_auth_failures": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "3",
			ValidateFunc: validateInteger,
		},
		"lockout_period": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "900",
			ValidateFunc: validateInteger,
		},
		"min_length": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "15",
			ValidateFunc: validateInteger,
		},
		"min_lowercase": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"min_numeric": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"min_special": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"min_uppercase": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"password_expiration_days": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "365",
			ValidateFunc: validateInteger,
		},
		"password_history": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "5",
			ValidateFunc: validateInteger,
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

func resourceAviPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPasswordPolicyCreate,
		Read:   ResourceAviPasswordPolicyRead,
		Update: resourceAviPasswordPolicyUpdate,
		Delete: resourceAviPasswordPolicyDelete,
		Schema: ResourcePasswordPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourcePasswordPolicyImporter,
		},
	}
}

func ResourcePasswordPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourcePasswordPolicySchema()
	return ResourceImporter(d, m, "passwordpolicy", s)
}

func ResourceAviPasswordPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePasswordPolicySchema()
	err := APIRead(d, meta, "passwordpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviPasswordPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePasswordPolicySchema()
	err := APICreate(d, meta, "passwordpolicy", s)
	if err == nil {
		err = ResourceAviPasswordPolicyRead(d, meta)
	}
	return err
}

func resourceAviPasswordPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePasswordPolicySchema()
	var err error
	err = APIUpdate(d, meta, "passwordpolicy", s)
	if err == nil {
		err = ResourceAviPasswordPolicyRead(d, meta)
	}
	return err
}

func resourceAviPasswordPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "passwordpolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
