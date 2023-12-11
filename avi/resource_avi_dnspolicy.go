// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

//nolint
func ResourceDnsPolicySchema() map[string]*schema.Schema {
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
		"internal": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
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
		"rule": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceDnsRuleSchema(),
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

//nolint
func resourceAviDnsPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviDnsPolicyCreate,
		Read:   ResourceAviDnsPolicyRead,
		Update: resourceAviDnsPolicyUpdate,
		Delete: resourceAviDnsPolicyDelete,
		Schema: ResourceDnsPolicySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceDnsPolicyImporter,
		},
	}
}

//nolint
func ResourceDnsPolicyImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceDnsPolicySchema()
	return ResourceImporter(d, m, "dnspolicy", s)
}

//nolint
func ResourceAviDnsPolicyRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := APIRead(d, meta, "dnspolicy", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

//nolint
func resourceAviDnsPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	err := APICreateOrUpdate(d, meta, "dnspolicy", s)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	return err
}

//nolint
func resourceAviDnsPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceDnsPolicySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "dnspolicy", s)
	if err == nil {
		err = ResourceAviDnsPolicyRead(d, meta)
	}
	return err
}

//nolint
func resourceAviDnsPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "dnspolicy")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
