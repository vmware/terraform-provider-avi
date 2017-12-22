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

func ResourceAlertEmailConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cc_emails": &schema.Schema{
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
		"to_emails": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviAlertEmailConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAlertEmailConfigCreate,
		Read:   ResourceAviAlertEmailConfigRead,
		Update: resourceAviAlertEmailConfigUpdate,
		Delete: resourceAviAlertEmailConfigDelete,
		Schema: ResourceAlertEmailConfigSchema(),
	}
}

func ResourceAviAlertEmailConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertEmailConfigSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/alertemailconfig/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
}

func resourceAviAlertEmailConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertEmailConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "alertemailconfig", s)
	if err == nil {
		err = ResourceAviAlertEmailConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertEmailConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertEmailConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "alertemailconfig", s)
	if err == nil {
		err = ResourceAviAlertEmailConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertEmailConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "alertemailconfig"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviAlertEmailConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
