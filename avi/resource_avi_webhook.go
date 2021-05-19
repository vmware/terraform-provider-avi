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

func ResourceWebhookSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"callback_url": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"verification_token": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviWebhook() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWebhookCreate,
		Read:   ResourceAviWebhookRead,
		Update: resourceAviWebhookUpdate,
		Delete: resourceAviWebhookDelete,
		Schema: ResourceWebhookSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceWebhookImporter,
		},
	}
}

func ResourceWebhookImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWebhookSchema()
	return ResourceImporter(d, m, "webhook", s)
}

func ResourceAviWebhookRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	err := APIRead(d, meta, "webhook", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	err := APICreateOrUpdate(d, meta, "webhook", s)
	if err == nil {
		err = ResourceAviWebhookRead(d, meta)
	}
	return err
}

func resourceAviWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "webhook", s)
	if err == nil {
		err = ResourceAviWebhookRead(d, meta)
	}
	return err
}

func resourceAviWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "webhook"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviWebhookDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
