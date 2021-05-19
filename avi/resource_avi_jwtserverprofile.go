/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceJWTServerProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"issuer": {
			Type:     schema.TypeString,
			Required: true,
		},
		"jwks_keys": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviJWTServerProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviJWTServerProfileCreate,
		Read:   ResourceAviJWTServerProfileRead,
		Update: resourceAviJWTServerProfileUpdate,
		Delete: resourceAviJWTServerProfileDelete,
		Schema: ResourceJWTServerProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceJWTServerProfileImporter,
		},
	}
}

func ResourceJWTServerProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceJWTServerProfileSchema()
	return ResourceImporter(d, m, "jwtserverprofile", s)
}

func ResourceAviJWTServerProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceJWTServerProfileSchema()
	err := APIRead(d, meta, "jwtserverprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviJWTServerProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceJWTServerProfileSchema()
	err := APICreateOrUpdate(d, meta, "jwtserverprofile", s)
	if err == nil {
		err = ResourceAviJWTServerProfileRead(d, meta)
	}
	return err
}

func resourceAviJWTServerProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceJWTServerProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "jwtserverprofile", s)
	if err == nil {
		err = ResourceAviJWTServerProfileRead(d, meta)
	}
	return err
}

func resourceAviJWTServerProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "jwtserverprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviJWTServerProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
