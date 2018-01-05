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

func ResourceWebhookSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"callback_url": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"verification_token": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
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
	}
}

func ResourceAviWebhookRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	err := ApiRead(d, meta, "webhook", s)
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
	err := ApiCreateOrUpdate(d, meta, "webhook", s)
	if err == nil {
		err = ResourceAviWebhookRead(d, meta)
	}
	return err
}

func resourceAviWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "webhook"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviWebhookDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
