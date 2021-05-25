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

func ResourceJWTProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"is_federated": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"jwks_keys": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceJWSKeySchema(),
		},
		"jwt_auth_type": {
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

func resourceAviJWTProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviJWTProfileCreate,
		Read:   ResourceAviJWTProfileRead,
		Update: resourceAviJWTProfileUpdate,
		Delete: resourceAviJWTProfileDelete,
		Schema: ResourceJWTProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceJWTProfileImporter,
		},
	}
}

func ResourceJWTProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceJWTProfileSchema()
	return ResourceImporter(d, m, "jwtprofile", s)
}

func ResourceAviJWTProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceJWTProfileSchema()
	err := APIRead(d, meta, "jwtprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviJWTProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceJWTProfileSchema()
	err := APICreateOrUpdate(d, meta, "jwtprofile", s)
	if err == nil {
		err = ResourceAviJWTProfileRead(d, meta)
	}
	return err
}

func resourceAviJWTProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceJWTProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "jwtprofile", s)
	if err == nil {
		err = ResourceAviJWTProfileRead(d, meta)
	}
	return err
}

func resourceAviJWTProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "jwtprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviJWTProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
