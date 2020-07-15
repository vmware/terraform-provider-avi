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

func ResourceTenantSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"config_settings": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTenantConfigurationSchema(),
		},
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
		"local": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"name": {
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

func resourceAviTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTenantCreate,
		Read:   ResourceAviTenantRead,
		Update: resourceAviTenantUpdate,
		Delete: resourceAviTenantDelete,
		Schema: ResourceTenantSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTenantImporter,
		},
	}
}

func ResourceTenantImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTenantSchema()
	return ResourceImporter(d, m, "tenant", s)
}

func ResourceAviTenantRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := ApiRead(d, meta, "tenant", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTenantCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := ApiCreateOrUpdate(d, meta, "tenant", s)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	return err
}

func resourceAviTenantUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "tenant", s)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	return err
}

func resourceAviTenantDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "tenant"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviTenantDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
