// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTLSProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"certificate_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
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
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"pki_profile_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
			ForceNew: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviTLSProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTLSProfileCreate,
		Read:   ResourceAviTLSProfileRead,
		Update: resourceAviTLSProfileUpdate,
		Delete: resourceAviTLSProfileDelete,
		Schema: ResourceTLSProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTLSProfileImporter,
		},
	}
}

func ResourceTLSProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTLSProfileSchema()
	return ResourceImporter(d, m, "tlsprofile", s)
}

func ResourceAviTLSProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTLSProfileSchema()
	err := APIRead(d, meta, "tlsprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTLSProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTLSProfileSchema()
	err := APICreate(d, meta, "tlsprofile", s)
	if err == nil {
		err = ResourceAviTLSProfileRead(d, meta)
	}
	return err
}

func resourceAviTLSProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTLSProfileSchema()
	var err error
	err = APIUpdate(d, meta, "tlsprofile", s)
	if err == nil {
		err = ResourceAviTLSProfileRead(d, meta)
	}
	return err
}

func resourceAviTLSProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "tlsprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
