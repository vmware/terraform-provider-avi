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

func ResourceMicroServiceGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"created_by": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"service_refs": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceAviMicroServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviMicroServiceGroupCreate,
		Read:   ResourceAviMicroServiceGroupRead,
		Update: resourceAviMicroServiceGroupUpdate,
		Delete: resourceAviMicroServiceGroupDelete,
		Schema: ResourceMicroServiceGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceMicroServiceGroupImporter,
		},
	}
}

func ResourceMicroServiceGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceMicroServiceGroupSchema()
	return ResourceImporter(d, m, "microservicegroup", s)
}

func ResourceAviMicroServiceGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceMicroServiceGroupSchema()
	err := ApiRead(d, meta, "microservicegroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviMicroServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceMicroServiceGroupSchema()
	err := ApiCreateOrUpdate(d, meta, "microservicegroup", s)
	if err == nil {
		err = ResourceAviMicroServiceGroupRead(d, meta)
	}
	return err
}

func resourceAviMicroServiceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceMicroServiceGroupSchema()
	var err error

	err = ApiCreateOrUpdate(d, meta, "microservicegroup", s)
	if err == nil {
		err = ResourceAviMicroServiceGroupRead(d, meta)
	}
	return err
}

func resourceAviMicroServiceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "microservicegroup"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviMicroServiceGroupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
