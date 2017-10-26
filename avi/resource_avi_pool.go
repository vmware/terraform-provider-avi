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
)

func ResourcePoolSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"url": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"health_monitor_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"servers": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceServer(),
		},
		"fail_action": &schema.Schema{
			Type:        schema.TypeSet,
			Description: "Fail action",
			Optional:    true,
			Set: func(v interface{}) int {
				return 0
			},
			Elem: ResourceFailAction(),
		},
	}
}

func resourceAviPool() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviPoolCreate,
		Read:   ResourceAviPoolRead,
		Update: resourceAviPoolUpdate,
		Delete: resourceAviPoolDelete,
		Schema: ResourcePoolSchema(),
	}
}

func ResourceAviPoolRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	log.Printf("[INFO] ResourceAviPoolRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/pool/" + uuid.(string)
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
	log.Printf("ResourceAviPoolRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviPoolRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviPoolRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviPoolRead Updated %v\n", d)
	return nil
}

func resourceAviPoolUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	err := ApiCreateOrUpdate(d, meta, "pool", s)
	if err == nil {
		err = ResourceAviPoolRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "pool", d)
	return err
}

func resourceAviPoolCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourcePoolSchema()
	err := ApiCreateOrUpdate(d, meta, "pool", s)
	log.Printf("[DEBUG] created object %v: %v", "pool", d)
	if err == nil {
		err = ResourceAviPoolRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "pool", d)
	return err
}

func resourceAviPoolDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "pool"
	log.Println("[INFO] ResourceAviPoolRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil {
			log.Println("[INFO] resourceAviPoolDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
