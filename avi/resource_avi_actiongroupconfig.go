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

func ResourceActionGroupConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_script_config_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"autoscale_trigger_notification": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"email_config_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"external_only": &schema.Schema{
			Type:     schema.TypeBool,
			Required: true,
		},
		"level": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"snmp_trap_profile_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"syslog_config_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
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

func resourceAviActionGroupConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviActionGroupConfigCreate,
		Read:   ResourceAviActionGroupConfigRead,
		Update: resourceAviActionGroupConfigUpdate,
		Delete: resourceAviActionGroupConfigDelete,
		Schema: ResourceActionGroupConfigSchema(),
	}
}

func ResourceAviActionGroupConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceActionGroupConfigSchema()
	log.Printf("[INFO] ResourceAviActionGroupConfigRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/actiongroupconfig/" + uuid.(string)
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
	log.Printf("ResourceAviActionGroupConfigRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviActionGroupConfigRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviActionGroupConfigRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviActionGroupConfigRead Updated %v\n", d)
	return nil
}

func resourceAviActionGroupConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceActionGroupConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "actiongroupconfig", s)
	log.Printf("[DEBUG] created object %v: %v", "actiongroupconfig", d)
	if err == nil {
		err = ResourceAviActionGroupConfigRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "actiongroupconfig", d)
	return err
}

func resourceAviActionGroupConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceActionGroupConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "actiongroupconfig", s)
	if err == nil {
		err = ResourceAviActionGroupConfigRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "actiongroupconfig", d)
	return err
}

func resourceAviActionGroupConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "actiongroupconfig"
	log.Println("[INFO] ResourceAviActionGroupConfigRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviActionGroupConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
