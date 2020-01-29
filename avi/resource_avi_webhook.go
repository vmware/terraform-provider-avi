/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
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
	err := ApiRead(d, meta, "webhook", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	err := ApiCreateOrUpdate(d, meta, "webhook", s)
	if err == nil {
		err = ResourceAviWebhookRead(d, meta)
	}
	return err
}

func resourceAviWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "webhook", s)
	if err == nil {
		err = ResourceAviWebhookRead(d, meta)
	}
	return err
}

func resourceAviWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "webhook"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
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
