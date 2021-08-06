// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
)

func ResourceGslbGeoDbProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"entries": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceGslbGeoDbEntrySchema(),
		},
		"is_federated": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
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

func resourceAviGslbGeoDbProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviGslbGeoDbProfileCreate,
		Read:   ResourceAviGslbGeoDbProfileRead,
		Update: resourceAviGslbGeoDbProfileUpdate,
		Delete: resourceAviGslbGeoDbProfileDelete,
		Schema: ResourceGslbGeoDbProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceGslbGeoDbProfileImporter,
		},
	}
}

func ResourceGslbGeoDbProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceGslbGeoDbProfileSchema()
	return ResourceImporter(d, m, "gslbgeodbprofile", s)
}

func ResourceAviGslbGeoDbProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbGeoDbProfileSchema()
	err := APIRead(d, meta, "gslbgeodbprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviGslbGeoDbProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbGeoDbProfileSchema()
	err := APICreateOrUpdate(d, meta, "gslbgeodbprofile", s)
	if err == nil {
		err = ResourceAviGslbGeoDbProfileRead(d, meta)
	}
	return err
}

func resourceAviGslbGeoDbProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceGslbGeoDbProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "gslbgeodbprofile", s)
	if err == nil {
		err = ResourceAviGslbGeoDbProfileRead(d, meta)
	}
	return err
}

func resourceAviGslbGeoDbProfileDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "gslbgeodbprofile"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviGslbGeoDbProfileDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
