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

func ResourceCustomerPortalInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"polling_interval": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  10,
		},
		"portal_url": {
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

func resourceAviCustomerPortalInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviCustomerPortalInfoCreate,
		Read:   ResourceAviCustomerPortalInfoRead,
		Update: resourceAviCustomerPortalInfoUpdate,
		Delete: resourceAviCustomerPortalInfoDelete,
		Schema: ResourceCustomerPortalInfoSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceCustomerPortalInfoImporter,
		},
	}
}

func ResourceCustomerPortalInfoImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceCustomerPortalInfoSchema()
	return ResourceImporter(d, m, "customerportalinfo", s)
}

func ResourceAviCustomerPortalInfoRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomerPortalInfoSchema()
	err := ApiRead(d, meta, "customerportalinfo", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviCustomerPortalInfoCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomerPortalInfoSchema()
	err := ApiCreateOrUpdate(d, meta, "customerportalinfo", s)
	if err == nil {
		err = ResourceAviCustomerPortalInfoRead(d, meta)
	}
	return err
}

func resourceAviCustomerPortalInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceCustomerPortalInfoSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "customerportalinfo", s)
	if err == nil {
		err = ResourceAviCustomerPortalInfoRead(d, meta)
	}
	return err
}

func resourceAviCustomerPortalInfoDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "customerportalinfo"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviCustomerPortalInfoDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
