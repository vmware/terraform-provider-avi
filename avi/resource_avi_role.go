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
		"filters": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"privileges": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePermissionSchema(),
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

func resourceAviRole() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviRoleCreate,
		Read:   ResourceAviRoleRead,
		Update: resourceAviRoleUpdate,
		Delete: resourceAviRoleDelete,
		Schema: ResourceRoleSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceRoleImporter,
		},
	}
}

func ResourceRoleImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceRoleSchema()
	return ResourceImporter(d, m, "role", s)
}

func ResourceAviRoleRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	err := ApiRead(d, meta, "role", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviRoleCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	err := ApiCreateOrUpdate(d, meta, "role", s)
	if err == nil {
		err = ResourceAviRoleRead(d, meta)
	}
	return err
}

func resourceAviRoleUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceRoleSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "role", s)
	if err == nil {
		err = ResourceAviRoleRead(d, meta)
	}
	return err
}

func resourceAviRoleDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "role"
	client := meta.(*clients.AviClient)
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviRoleDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
