// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceBotMappingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mapping_rules": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceBotMappingRuleSchema(),
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
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviBotMapping() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviBotMappingCreate,
		Read:   ResourceAviBotMappingRead,
		Update: resourceAviBotMappingUpdate,
		Delete: resourceAviBotMappingDelete,
		Schema: ResourceBotMappingSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceBotMappingImporter,
		},
	}
}

func ResourceBotMappingImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceBotMappingSchema()
	return ResourceImporter(d, m, "botmapping", s)
}

func ResourceAviBotMappingRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotMappingSchema()
	err := APIRead(d, meta, "botmapping", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviBotMappingCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotMappingSchema()
	err := APICreate(d, meta, "botmapping", s)
	if err == nil {
		err = ResourceAviBotMappingRead(d, meta)
	}
	return err
}

func resourceAviBotMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotMappingSchema()
	var err error
	err = APIUpdate(d, meta, "botmapping", s)
	if err == nil {
		err = ResourceAviBotMappingRead(d, meta)
	}
	return err
}

func resourceAviBotMappingDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "botmapping")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
