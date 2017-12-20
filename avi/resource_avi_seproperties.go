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
	log.Printf("[INFO] ResourceAviSePropertiesRead Avi Client %v\n", d)
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
	// no need to set the ID
	log.Printf("ResourceAviSePropertiesRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviSePropertiesRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviSePropertiesRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviSePropertiesRead Updated %v\n", d)
	return nil
}

func resourceAviSePropertiesCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSePropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "seproperties", s)
	log.Printf("[DEBUG] created object %v: %v", "seproperties", d)
	if err == nil {
		err = ResourceAviSePropertiesRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "seproperties", d)
	return err
}

func resourceAviSePropertiesUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSePropertiesSchema()
	err := ApiCreateOrUpdate(d, meta, "seproperties", s)
	if err == nil {
		err = ResourceAviSePropertiesRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "seproperties", d)
	return err
}

func resourceAviSePropertiesDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "seproperties"
	log.Println("[INFO] ResourceAviSePropertiesRead Avi Client")
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
