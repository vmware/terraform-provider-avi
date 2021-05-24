// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceBotConfigConsolidatorSchema() map[string]*schema.Schema {
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
		"script": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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

func resourceAviBotConfigConsolidator() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviBotConfigConsolidatorCreate,
		Read:   ResourceAviBotConfigConsolidatorRead,
		Update: resourceAviBotConfigConsolidatorUpdate,
		Delete: resourceAviBotConfigConsolidatorDelete,
		Schema: ResourceBotConfigConsolidatorSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceBotConfigConsolidatorImporter,
		},
	}
}

func ResourceBotConfigConsolidatorImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceBotConfigConsolidatorSchema()
	return ResourceImporter(d, m, "botconfigconsolidator", s)
}

func ResourceAviBotConfigConsolidatorRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotConfigConsolidatorSchema()
	err := APIRead(d, meta, "botconfigconsolidator", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviBotConfigConsolidatorCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotConfigConsolidatorSchema()
	err := APICreateOrUpdate(d, meta, "botconfigconsolidator", s)
	if err == nil {
		err = ResourceAviBotConfigConsolidatorRead(d, meta)
	}
	return err
}

func resourceAviBotConfigConsolidatorUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotConfigConsolidatorSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "botconfigconsolidator", s)
	if err == nil {
		err = ResourceAviBotConfigConsolidatorRead(d, meta)
	}
	return err
}

func resourceAviBotConfigConsolidatorDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "botconfigconsolidator"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviBotConfigConsolidatorDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
