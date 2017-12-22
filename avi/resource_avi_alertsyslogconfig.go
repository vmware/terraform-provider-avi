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

func ResourceAlertSyslogConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"syslog_servers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceAlertSyslogServerSchema(),
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
	}
}

func resourceAviAlertSyslogConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAlertSyslogConfigCreate,
		Read:   ResourceAviAlertSyslogConfigRead,
		Update: resourceAviAlertSyslogConfigUpdate,
		Delete: resourceAviAlertSyslogConfigDelete,
		Schema: ResourceAlertSyslogConfigSchema(),
	}
}

func ResourceAviAlertSyslogConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertSyslogConfigSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/alertsyslogconfig/" + uuid.(string)
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

func resourceAviAlertSyslogConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertSyslogConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "alertsyslogconfig", s)
	if err == nil {
		err = ResourceAviAlertSyslogConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertSyslogConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertSyslogConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "alertsyslogconfig", s)
	if err == nil {
		err = ResourceAviAlertSyslogConfigRead(d, meta)
	}
	return err
}

func resourceAviAlertSyslogConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "alertsyslogconfig"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviAlertSyslogConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
