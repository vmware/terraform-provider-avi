// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceLabelGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"labels": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleMatchOperationMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviLabelGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviLabelGroupCreate,
		Read:   ResourceAviLabelGroupRead,
		Update: resourceAviLabelGroupUpdate,
		Delete: resourceAviLabelGroupDelete,
		Schema: ResourceLabelGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceLabelGroupImporter,
		},
	}
}

func ResourceLabelGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceLabelGroupSchema()
	return ResourceImporter(d, m, "labelgroup", s)
}

func ResourceAviLabelGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLabelGroupSchema()
	err := APIRead(d, meta, "labelgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviLabelGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLabelGroupSchema()
	err := APICreateOrUpdate(d, meta, "labelgroup", s)
	if err == nil {
		err = ResourceAviLabelGroupRead(d, meta)
	}
	return err
}

func resourceAviLabelGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLabelGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "labelgroup", s)
	if err == nil {
		err = ResourceAviLabelGroupRead(d, meta)
	}
	return err
}

func resourceAviLabelGroupDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "labelgroup")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
