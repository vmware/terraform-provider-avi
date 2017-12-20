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

func ResourceAlertScriptConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"action_script": &schema.Schema{
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
	}
}

func resourceAviAlertScriptConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAlertScriptConfigCreate,
		Read:   ResourceAviAlertScriptConfigRead,
		Update: resourceAviAlertScriptConfigUpdate,
		Delete: resourceAviAlertScriptConfigDelete,
		Schema: ResourceAlertScriptConfigSchema(),
	}
}

func ResourceAviAlertScriptConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertScriptConfigSchema()
	log.Printf("[INFO] ResourceAviAlertScriptConfigRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/alertscriptconfig/" + uuid.(string)
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
	log.Printf("ResourceAviAlertScriptConfigRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviAlertScriptConfigRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviAlertScriptConfigRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviAlertScriptConfigRead Updated %v\n", d)
	return nil
}

func resourceAviAlertScriptConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertScriptConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "alertscriptconfig", s)
	log.Printf("[DEBUG] created object %v: %v", "alertscriptconfig", d)
	if err == nil {
		err = ResourceAviAlertScriptConfigRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "alertscriptconfig", d)
	return err
}

func resourceAviAlertScriptConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAlertScriptConfigSchema()
	err := ApiCreateOrUpdate(d, meta, "alertscriptconfig", s)
	if err == nil {
		err = ResourceAviAlertScriptConfigRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "alertscriptconfig", d)
	return err
}

func resourceAviAlertScriptConfigDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "alertscriptconfig"
	log.Println("[INFO] ResourceAviAlertScriptConfigRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviAlertScriptConfigDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
