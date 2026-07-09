// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceLabelProfileSchema() map[string]*schema.Schema {
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
		"label_definitions": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviLabelProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviLabelProfileCreate,
		Read:   ResourceAviLabelProfileRead,
		Update: resourceAviLabelProfileUpdate,
		Delete: resourceAviLabelProfileDelete,
		Schema: ResourceLabelProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceLabelProfileImporter,
		},
	}
}

func ResourceLabelProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceLabelProfileSchema()
	return ResourceImporter(d, m, "labelprofile", s)
}

func ResourceAviLabelProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLabelProfileSchema()
	err := APIRead(d, meta, "labelprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviLabelProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLabelProfileSchema()
	err := APICreate(d, meta, "labelprofile", s)
	if err == nil {
		err = ResourceAviLabelProfileRead(d, meta)
	}
	return err
}

func resourceAviLabelProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLabelProfileSchema()
	var err error
	err = APIUpdate(d, meta, "labelprofile", s)
	if err == nil {
		err = ResourceAviLabelProfileRead(d, meta)
	}
	return err
}

func resourceAviLabelProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "labelprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
