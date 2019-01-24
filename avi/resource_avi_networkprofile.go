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

func ResourceNetworkProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"connection_mirror": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"description": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"profile": &schema.Schema{
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceNetworkProfileUnionSchema(),
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviNetworkProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkProfileCreate,
		Read:   ResourceAviNetworkProfileRead,
		Update: resourceAviNetworkProfileUpdate,
		Delete: resourceAviNetworkProfileDelete,
		Schema: ResourceNetworkProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkProfileImporter,
		},
	}
}

func ResourceNetworkProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkProfileSchema()
	return ResourceImporter(d, m, "networkprofile", s)
}

func ResourceAviNetworkProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkProfileSchema()
	err := ApiRead(d, meta, "networkprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkProfileSchema()
	err := ApiCreateOrUpdate(d, meta, "networkprofile", s)
	if err == nil {
		err = ResourceAviNetworkProfileRead(d, meta)
	}
	return err
}

func resourceAviNetworkProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkProfileSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "networkprofile", s)
	if err == nil {
		err = ResourceAviNetworkProfileRead(d, meta)
	}
	return err
}

func resourceAviNetworkProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "networkprofile"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviNetworkProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
