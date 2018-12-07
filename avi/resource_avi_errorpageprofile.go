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
	"time"
)

func ResourceErrorPageProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"error_pages": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceErrorPageSchema(),
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
		Importer: &schema.ResourceImporter{
			State: ResourceErrorPageProfileImporter,
		},
	}
}

func ResourceErrorPageProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceErrorPageProfileSchema()
	return ResourceImporter(d, m, "errorpageprofile", s)
}

func ResourceAviErrorPageProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	err := ApiRead(d, meta, "errorpageprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviErrorPageProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "errorpageprofile", s)
	if err == nil {
		err = ResourceAviErrorPageProfileRead(d, meta)
	}
	return err
}

func resourceAviErrorPageProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "errorpageprofile", s)
	if err == nil {
		err = ResourceAviErrorPageProfileRead(d, meta)
	}
	return err
}

func resourceAviErrorPageProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "errorpageprofile"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviErrorPageProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
