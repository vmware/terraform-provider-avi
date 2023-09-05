// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceStringGroupSchema() map[string]*schema.Schema {
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
		"kv": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"longest_match": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
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
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
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

func resourceAviStringGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviStringGroupCreate,
		Read:   ResourceAviStringGroupRead,
		Update: resourceAviStringGroupUpdate,
		Delete: resourceAviStringGroupDelete,
		Schema: ResourceStringGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceStringGroupImporter,
		},
	}
}

func ResourceStringGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceStringGroupSchema()
	return ResourceImporter(d, m, "stringgroup", s)
}

func ResourceAviStringGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStringGroupSchema()
	err := APIRead(d, meta, "stringgroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviStringGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStringGroupSchema()
	err := APICreateOrUpdate(d, meta, "stringgroup", s)
	if err == nil {
		err = ResourceAviStringGroupRead(d, meta)
	}
	return err
}

func resourceAviStringGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceStringGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "stringgroup", s)
	if err == nil {
		err = ResourceAviStringGroupRead(d, meta)
	}
	return err
}

func resourceAviStringGroupDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "stringgroup")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
