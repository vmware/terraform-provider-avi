/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
)

func ResourceControllerSiteSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"address": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
		},
		"port": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  443,
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

func resourceAviControllerSite() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviControllerSiteCreate,
		Read:   ResourceAviControllerSiteRead,
		Update: resourceAviControllerSiteUpdate,
		Delete: resourceAviControllerSiteDelete,
		Schema: ResourceControllerSiteSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceControllerSiteImporter,
		},
	}
}

func ResourceControllerSiteImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceControllerSiteSchema()
	return ResourceImporter(d, m, "controllersite", s)
}

func ResourceAviControllerSiteRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerSiteSchema()
	err := ApiRead(d, meta, "controllersite", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviControllerSiteCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerSiteSchema()
	err := ApiCreateOrUpdate(d, meta, "controllersite", s)
	if err == nil {
		err = ResourceAviControllerSiteRead(d, meta)
	}
	return err
}

func resourceAviControllerSiteUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceControllerSiteSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "controllersite", s)
	if err == nil {
		err = ResourceAviControllerSiteRead(d, meta)
	}
	return err
}

func resourceAviControllerSiteDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "controllersite"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviControllerSiteDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
