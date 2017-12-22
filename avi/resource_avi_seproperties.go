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

func ResourceSePropertiesSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"se_agent_properties": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSeAgentPropertiesSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"se_bootup_properties": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSeBootupPropertiesSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"se_runtime_properties": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceSeRuntimePropertiesSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSeProperties() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSePropertiesCreate,
		Read:   ResourceAviSePropertiesRead,
		Update: resourceAviSePropertiesUpdate,
		Delete: resourceAviSePropertiesDelete,
		Schema: ResourceSePropertiesSchema(),
	}
}

func ResourceAviSePropertiesRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSePropertiesSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/seproperties/" + uuid.(string)
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

func resourceAviSePropertiesCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSePropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "seproperties", s)
	if err == nil {
		err = ResourceAviSePropertiesRead(d, meta)
	}
	return err
}

func resourceAviSePropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSePropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "seproperties", s)
	if err == nil {
		err = ResourceAviSePropertiesRead(d, meta)
	}
	return err
}

func resourceAviSePropertiesDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "seproperties"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviSePropertiesDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
