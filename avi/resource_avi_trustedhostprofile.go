// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTrustedHostProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"hosts": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     ResourceTrustedHostSchema(),
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

func resourceAviTrustedHostProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTrustedHostProfileCreate,
		Read:   ResourceAviTrustedHostProfileRead,
		Update: resourceAviTrustedHostProfileUpdate,
		Delete: resourceAviTrustedHostProfileDelete,
		Schema: ResourceTrustedHostProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTrustedHostProfileImporter,
		},
	}
}

func ResourceTrustedHostProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTrustedHostProfileSchema()
	return ResourceImporter(d, m, "trustedhostprofile", s)
}

func ResourceAviTrustedHostProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrustedHostProfileSchema()
	err := APIRead(d, meta, "trustedhostprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTrustedHostProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrustedHostProfileSchema()
	err := APICreate(d, meta, "trustedhostprofile", s)
	if err == nil {
		err = ResourceAviTrustedHostProfileRead(d, meta)
	}
	return err
}

func resourceAviTrustedHostProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTrustedHostProfileSchema()
	var err error
	err = APIUpdate(d, meta, "trustedhostprofile", s)
	if err == nil {
		err = ResourceAviTrustedHostProfileRead(d, meta)
	}
	return err
}

func resourceAviTrustedHostProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "trustedhostprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
