/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceBotMappingSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"mapping_rules": {
			Type:     schema.TypeList,
			Required: true,
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
	err := APICreateOrUpdate(d, meta, "botmapping", s)
	if err == nil {
		err = ResourceAviBotMappingRead(d, meta)
	}
	return err
}

func resourceAviBotMappingUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBotMappingSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "botmapping", s)
	if err == nil {
		err = ResourceAviBotMappingRead(d, meta)
	}
	return err
}

func resourceAviBotMappingDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "botmapping"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviBotMappingDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
