// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceWafCRSSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
		},
		"integrity": {
			Type:     schema.TypeString,
			Required: true,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"release_date": {
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
		"version": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourceAviWafCRS() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafCRSCreate,
		Read:   ResourceAviWafCRSRead,
		Update: resourceAviWafCRSUpdate,
		Delete: resourceAviWafCRSDelete,
		Schema: ResourceWafCRSSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafCRSImporter,
		},
	}
}

func ResourceWafCRSImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafCRSSchema()
	return ResourceImporter(d, m, "wafcrs", s)
}

func ResourceAviWafCRSRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafCRSSchema()
	err := APIRead(d, meta, "wafcrs", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafCRSCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafCRSSchema()
	err := APICreateOrUpdate(d, meta, "wafcrs", s)
	if err == nil {
		err = ResourceAviWafCRSRead(d, meta)
	}
	return err
}

func resourceAviWafCRSUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafCRSSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "wafcrs", s)
	if err == nil {
		err = ResourceAviWafCRSRead(d, meta)
	}
	return err
}

func resourceAviWafCRSDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "wafcrs")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
