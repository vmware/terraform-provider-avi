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

func ResourceAlertEmailConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cc_emails": {
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
		"to_emails": {
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

func resourceAviAlertEmailConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAlertEmailConfigCreate,
		Read:   ResourceAviAlertEmailConfigRead,
		Update: resourceAviAlertEmailConfigUpdate,
		Delete: resourceAviAlertEmailConfigDelete,
		Schema: ResourceAlertEmailConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAlertEmailConfigImporter,
		},
	}
}

func ResourceAlertEmailConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAlertEmailConfigSchema()
	return ResourceImporter(d, m, "alertemailconfig", s)
}

func ResourceAviAlertEmailConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertEmailConfigSchema()
	err := ApiRead(d, meta, "alertemailconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
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
	var err error
	err = ApiCreateOrUpdate(d, meta, "alertemailconfig", s)
	if err == nil {
		err = ResourceAviAlertEmailConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertEmailConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "alertemailconfig"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviAlertEmailConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
