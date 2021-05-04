/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import (
	"log"
	"strings"

	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceImageSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"controller_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePackageDetailsSchema(),
		},
		"controller_patch_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"controller_patch_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"migrations": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSupportedMigrationsSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"se_info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourcePackageDetailsSchema(),
		},
		"se_patch_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_patch_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uber_bundle": {
			Type:     schema.TypeBool,
			Optional: true,
			Default:  false,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviImage() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviImageCreate,
		Read:   ResourceAviImageRead,
		Update: resourceAviImageUpdate,
		Delete: resourceAviImageDelete,
		Schema: ResourceImageSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceImageImporter,
		},
	}
}

func ResourceImageImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceImageSchema()
	return ResourceImporter(d, m, "image", s)
}

func ResourceAviImageRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceImageSchema()
	err := APIRead(d, meta, "image", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviImageCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceImageSchema()
	err := APICreateOrUpdate(d, meta, "image", s)
	if err == nil {
		err = ResourceAviImageRead(d, meta)
	}
	return err
}

func resourceAviImageUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceImageSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "image", s)
	if err == nil {
		err = ResourceAviImageRead(d, meta)
	}
	return err
}

func resourceAviImageDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "image"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviImageDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
