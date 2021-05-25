/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceL4PolicySetSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"is_internal_policy": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"l4_connection_policy": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceL4ConnectionPolicySchema(),
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
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviL4PolicySet() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviL4PolicySetCreate,
		Read:   ResourceAviL4PolicySetRead,
		Update: resourceAviL4PolicySetUpdate,
		Delete: resourceAviL4PolicySetDelete,
		Schema: ResourceL4PolicySetSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceL4PolicySetImporter,
		},
	}
}

func ResourceL4PolicySetImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceL4PolicySetSchema()
	return ResourceImporter(d, m, "l4policyset", s)
}

func ResourceAviL4PolicySetRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceL4PolicySetSchema()
	err := APIRead(d, meta, "l4policyset", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviL4PolicySetCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceL4PolicySetSchema()
	err := APICreateOrUpdate(d, meta, "l4policyset", s)
	if err == nil {
		err = ResourceAviL4PolicySetRead(d, meta)
	}
	return err
}

func resourceAviL4PolicySetUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceL4PolicySetSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "l4policyset", s)
	if err == nil {
		err = ResourceAviL4PolicySetRead(d, meta)
	}
	return err
}

func resourceAviL4PolicySetDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "l4policyset"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviL4PolicySetDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
