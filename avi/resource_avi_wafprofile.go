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

func ResourceWafProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"config": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceWafConfigSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"files": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceWafDataFileSchema(),
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

func resourceAviWafProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviWafProfileCreate,
		Read:   ResourceAviWafProfileRead,
		Update: resourceAviWafProfileUpdate,
		Delete: resourceAviWafProfileDelete,
		Schema: ResourceWafProfileSchema(),
	}
}

func ResourceAviWafProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafProfileSchema()
	log.Printf("[INFO] ResourceAviWafProfileRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/wafprofile/" + uuid.(string)
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
	log.Printf("ResourceAviWafProfileRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviWafProfileRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviWafProfileRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviWafProfileRead Updated %v\n", d)
	return nil
}

func resourceAviWafProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "wafprofile", s)
	log.Printf("[DEBUG] created object %v: %v", "wafprofile", d)
	if err == nil {
		err = ResourceAviWafProfileRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "wafprofile", d)
	return err
}

func resourceAviWafProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "wafprofile", s)
	if err == nil {
		err = ResourceAviWafProfileRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "wafprofile", d)
	return err
}

func resourceAviWafProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafprofile"
	log.Println("[INFO] ResourceAviWafProfileRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviWafProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
