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

func ResourceErrorPageProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"app_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "VS Name",
		},
		"company_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "Avi Networks",
		},
		"error_pages": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceErrorPageSchema(),
		},
		"host_name": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "Host Header",
		},
		"name": &schema.Schema{
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

func resourceAviErrorPageProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviErrorPageProfileCreate,
		Read:   ResourceAviErrorPageProfileRead,
		Update: resourceAviErrorPageProfileUpdate,
		Delete: resourceAviErrorPageProfileDelete,
		Schema: ResourceErrorPageProfileSchema(),
	}
}

func ResourceAviErrorPageProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	log.Printf("[INFO] ResourceAviErrorPageProfileRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/errorpageprofile/" + uuid.(string)
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
	log.Printf("ResourceAviErrorPageProfileRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviErrorPageProfileRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviErrorPageProfileRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviErrorPageProfileRead Updated %v\n", d)
	return nil
}

func resourceAviErrorPageProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "errorpageprofile", s)
	log.Printf("[DEBUG] created object %v: %v", "errorpageprofile", d)
	if err == nil {
		err = ResourceAviErrorPageProfileRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "errorpageprofile", d)
	return err
}

func resourceAviErrorPageProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "errorpageprofile", s)
	if err == nil {
		err = ResourceAviErrorPageProfileRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "errorpageprofile", d)
	return err
}

func resourceAviErrorPageProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "errorpageprofile"
	log.Println("[INFO] ResourceAviErrorPageProfileRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviErrorPageProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
