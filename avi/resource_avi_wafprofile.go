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
		Importer: &schema.ResourceImporter{
			State: ResourceWafProfileImporter,
		},
	}
}

func ResourceWafProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceWafProfileSchema()
	return ResourceImporter(d, m, "wafprofile", s)
}

func ResourceAviWafProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafProfileSchema()
	err := ApiRead(d, meta, "wafprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviWafProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "wafprofile", s)
	if err == nil {
		err = ResourceAviWafProfileRead(d, meta)
	}
	return err
}

func resourceAviWafProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceWafProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "wafprofile", s)
	if err == nil {
		err = ResourceAviWafProfileRead(d, meta)
	}
	return err
}

func resourceAviWafProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "wafprofile"
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
