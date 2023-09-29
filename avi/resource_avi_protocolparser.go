// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceProtocolParserSchema() map[string]*schema.Schema {
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
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"parser_code": {
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

func resourceAviProtocolParser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviProtocolParserCreate,
		Read:   ResourceAviProtocolParserRead,
		Update: resourceAviProtocolParserUpdate,
		Delete: resourceAviProtocolParserDelete,
		Schema: ResourceProtocolParserSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceProtocolParserImporter,
		},
	}
}

func ResourceProtocolParserImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceProtocolParserSchema()
	return ResourceImporter(d, m, "protocolparser", s)
}

func ResourceAviProtocolParserRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceProtocolParserSchema()
	err := APIRead(d, meta, "protocolparser", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviProtocolParserCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceProtocolParserSchema()
	err := APICreateOrUpdate(d, meta, "protocolparser", s)
	if err == nil {
		err = ResourceAviProtocolParserRead(d, meta)
	}
	return err
}

func resourceAviProtocolParserUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceProtocolParserSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "protocolparser", s)
	if err == nil {
		err = ResourceAviProtocolParserRead(d, meta)
	}
	return err
}

func resourceAviProtocolParserDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "protocolparser")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
