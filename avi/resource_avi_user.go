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

func ResourceUserSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"access": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceUserRoleSchema(),
		},
		"default_tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"email": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"full_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"is_superuser": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"local": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"password": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"user_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"username": {
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

func resourceAviUser() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUserCreate,
		Read:   ResourceAviUserRead,
		Update: resourceAviUserUpdate,
		Delete: resourceAviUserDelete,
		Schema: ResourceUserSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceUserImporter,
		},
	}
}

func ResourceUserImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUserSchema()
	return ResourceImporter(d, m, "user", s)
}

func ResourceAviUserRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserSchema()
	err := APIRead(d, meta, "user", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUserCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserSchema()
	err := APICreateOrUpdate(d, meta, "user", s)
	if err == nil {
		err = ResourceAviUserRead(d, meta)
	}
	return err
}

func resourceAviUserUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUserSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "user", s)
	if err == nil {
		err = ResourceAviUserRead(d, meta)
	}
	return err
}

func resourceAviUserDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "user"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviUserDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
