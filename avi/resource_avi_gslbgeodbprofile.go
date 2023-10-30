// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceGslbGeoDbProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
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
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "gslbgeodbprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
