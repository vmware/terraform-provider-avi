// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceRetentionPolicySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"history": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRetentionSummarySchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"policy": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourcePolicySpecSchema(),
		},
		"summary": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRetentionSummarySchema(),
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

func resourceAviRetentionPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviRetentionPolicyCreate,
		Read:   ResourceAviRetentionPolicyRead,
		Update: resourceAviRetentionPolicyUpdate,
		Delete: resourceAviRetentionPolicyDelete,
		Schema: ResourceRetentionPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceRetentionPolicyImporter,
		},
	}
}

func ResourceRetentionPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceRetentionPolicySchema()
	return ResourceImporter(d, m, "retentionpolicy", s)
}

func ResourceAviRetentionPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRetentionPolicySchema()
	err := APIRead(d, meta, "retentionpolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviRetentionPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRetentionPolicySchema()
	err := APICreate(d, meta, "retentionpolicy", s)
	if err == nil {
		err = ResourceAviRetentionPolicyRead(d, meta)
	}
	return err
}

func resourceAviRetentionPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRetentionPolicySchema()
	var err error
	err = APIUpdate(d, meta, "retentionpolicy", s)
	if err == nil {
		err = ResourceAviRetentionPolicyRead(d, meta)
	}
	return err
}

func resourceAviRetentionPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "retentionpolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
