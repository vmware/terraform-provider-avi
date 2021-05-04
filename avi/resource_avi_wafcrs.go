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

func ResourceWafCRSSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Required: true,
		},
		"groups": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafRuleGroupSchema(),
		},
		"integrity": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"release_date": {
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
		"version": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourceAviWafCRS() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafCRSCreate,
		Read:   ResourceAviWafCRSRead,
		Update: resourceAviWafCRSUpdate,
		Delete: resourceAviWafCRSDelete,
		Schema: ResourceWafCRSSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWafCRSImporter,
		},
	}
}

func ResourceWafCRSImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafCRSSchema()
	return ResourceImporter(d, m, "wafcrs", s)
}

func ResourceAviWafCRSRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafCRSSchema()
	err := APIRead(d, meta, "wafcrs", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafCRSCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafCRSSchema()
	err := APICreateOrUpdate(d, meta, "wafcrs", s)
	if err == nil {
		err = ResourceAviWafCRSRead(d, meta)
	}
	return err
}

func resourceAviWafCRSUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafCRSSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "wafcrs", s)
	if err == nil {
		err = ResourceAviWafCRSRead(d, meta)
	}
	return err
}

func resourceAviWafCRSDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafcrs"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWafCRSDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
