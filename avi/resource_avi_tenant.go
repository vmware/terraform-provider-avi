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

func ResourceTenantSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"config_settings": &schema.Schema{
			Type:     schema.TypeSet,
			Optional: true,
			Elem:     ResourceTenantConfigurationSchema(),
			Set: func(v interface{}) int {
				return 0
			},
		},
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"local": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTenantCreate,
		Read:   ResourceAviTenantRead,
		Update: resourceAviTenantUpdate,
		Delete: resourceAviTenantDelete,
		Schema: ResourceTenantSchema(),
	}
}

func ResourceAviTenantRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	log.Printf("[INFO] ResourceAviTenantRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/tenant/" + uuid.(string)
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
	log.Printf("ResourceAviTenantRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviTenantRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviTenantRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviTenantRead Updated %v\n", d)
	return nil
}

func resourceAviTenantCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := ApiCreateOrUpdate(d, meta, "tenant", s)
	log.Printf("[DEBUG] created object %v: %v", "tenant", d)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "tenant", d)
	return err
}

func resourceAviTenantUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTenantSchema()
	err := ApiCreateOrUpdate(d, meta, "tenant", s)
	if err == nil {
		err = ResourceAviTenantRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "tenant", d)
	return err
}

func resourceAviTenantDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "tenant"
	log.Println("[INFO] ResourceAviTenantRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviTenantDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
