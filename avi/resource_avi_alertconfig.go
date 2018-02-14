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

func ResourceAlertConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_group_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"alert_rule": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true, Elem: ResourceAlertRuleSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"autoscale_alert": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"category": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"enabled": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"expiry_time": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  86400,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"obj_uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"object_type": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"recommendation": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"rolling_window": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},
		"source": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"summary": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"threshold": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  1,
		},
		"throttle": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  600,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviAlertConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAlertConfigCreate,
		Read:   ResourceAviAlertConfigRead,
		Update: resourceAviAlertConfigUpdate,
		Delete: resourceAviAlertConfigDelete,
		Schema: ResourceAlertConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAlertConfigImporter,
		},
	}
}

func ResourceAlertConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAlertConfigSchema()
	return ResourceImporter(d, m, "alertconfig", s)
}

func ResourceAviAlertConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertConfigSchema()
	err := ApiRead(d, meta, "alertconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAlertConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "alertconfig", s)
	if err == nil {
		err = ResourceAviAlertConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertConfigSchema()
	var err error

	err = ApiCreateOrUpdate(d, meta, "alertconfig", s)
	if err == nil {
		err = ResourceAviAlertConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "alertconfig"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviAlertConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
