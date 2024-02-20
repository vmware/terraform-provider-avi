// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSePropertiesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"se_agent_properties": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSeAgentPropertiesSchema(),
		},
		"se_bootup_properties": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSeBootupPropertiesSchema(),
		},
		"se_runtime_properties": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSeRuntimePropertiesSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSeProperties() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSePropertiesCreate,
		Read:   ResourceAviSePropertiesRead,
		Update: resourceAviSePropertiesUpdate,
		Delete: resourceAviSePropertiesDelete,
		Schema: ResourceSePropertiesSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSePropertiesImporter,
		},
	}
}

func ResourceSePropertiesImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSePropertiesSchema()
	return ResourceImporter(d, m, "seproperties", s)
}

func ResourceAviSePropertiesRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSePropertiesSchema()
	err := APIRead(d, meta, "seproperties", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSePropertiesCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSePropertiesSchema()
	err := APICreate(d, meta, "seproperties", s)
	if err == nil {
		err = ResourceAviSePropertiesRead(d, meta)
	}
	return err
}

func resourceAviSePropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSePropertiesSchema()
	var err error
	err = APIUpdate(d, meta, "seproperties", s)
	if err == nil {
		err = ResourceAviSePropertiesRead(d, meta)
	}
	return err
}

func resourceAviSePropertiesDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "seproperties")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
