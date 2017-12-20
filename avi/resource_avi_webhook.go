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
	log.Printf("[INFO] ResourceAviWebhookRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/webhook/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	// no need to set the ID
	log.Printf("ResourceAviWebhookRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviWebhookRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviWebhookRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviWebhookRead Updated %v\n", d)
	return nil
}

func resourceAviWebhookCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	err := ApiCreateOrUpdate(d, meta, "webhook", s)
	log.Printf("[DEBUG] created object %v: %v", "webhook", d)
	if err == nil {
		err = ResourceAviWebhookRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "webhook", d)
	return err
}

func resourceAviWebhookUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWebhookSchema()
	err := ApiCreateOrUpdate(d, meta, "webhook", s)
	if err == nil {
		err = ResourceAviWebhookRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "webhook", d)
	return err
}

func resourceAviWebhookDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "webhook"
	log.Println("[INFO] ResourceAviWebhookRead Avi Client")
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
