// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceProtocolParserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
	objType := "protocolparser"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviProtocolParserDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
