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

func ResourceRoleSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"privileges": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePermissionSchema(),
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

func resourceAviRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviRoleCreate,
		Read:   ResourceAviRoleRead,
		Update: resourceAviRoleUpdate,
		Delete: resourceAviRoleDelete,
		Schema: ResourceRoleSchema(),
	}
}

func ResourceAviRoleRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	log.Printf("[INFO] ResourceAviRoleRead Avi Client %v\n", d)
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/role/" + uuid.(string)
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
	log.Printf("ResourceAviRoleRead CURRENT obj %v\n", d)

	log.Printf("ResourceAviRoleRead Read API obj %v\n", obj)
	if tObj, err := ApiDataToSchema(obj, d, s); err == nil {
		log.Printf("[INFO] ResourceAviRoleRead Converted obj %v\n", tObj)
		//err = d.Set("obj", tObj)
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	log.Printf("[INFO] ResourceAviRoleRead Updated %v\n", d)
	return nil
}

func resourceAviRoleCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	err := ApiCreateOrUpdate(d, meta, "role", s)
	log.Printf("[DEBUG] created object %v: %v", "role", d)
	if err == nil {
		err = ResourceAviRoleRead(d, meta)
	}
	log.Printf("[DEBUG] created object %v: %v", "role", d)
	return err
}

func resourceAviRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	err := ApiCreateOrUpdate(d, meta, "role", s)
	if err == nil {
		err = ResourceAviRoleRead(d, meta)
	}
	log.Printf("[DEBUG] updated object %v: %v", "role", d)
	return err
}

func resourceAviRoleDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "role"
	log.Println("[INFO] ResourceAviRoleRead Avi Client")
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviRoleDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
