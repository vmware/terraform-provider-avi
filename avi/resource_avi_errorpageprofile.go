// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceErrorPageProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"error_pages": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceErrorPageSchema(),
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

func resourceAviErrorPageProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviErrorPageProfileCreate,
		Read:   ResourceAviErrorPageProfileRead,
		Update: resourceAviErrorPageProfileUpdate,
		Delete: resourceAviErrorPageProfileDelete,
		Schema: ResourceErrorPageProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceErrorPageProfileImporter,
		},
	}
}

func ResourceErrorPageProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceErrorPageProfileSchema()
	return ResourceImporter(d, m, "errorpageprofile", s)
}

func ResourceAviErrorPageProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	err := APIRead(d, meta, "errorpageprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviErrorPageProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	err := APICreateOrUpdate(d, meta, "errorpageprofile", s)
	if err == nil {
		err = ResourceAviErrorPageProfileRead(d, meta)
	}
	return err
}

func resourceAviErrorPageProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceErrorPageProfileSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "errorpageprofile", s)
	if err == nil {
		err = ResourceAviErrorPageProfileRead(d, meta)
	}
	return err
}

func resourceAviErrorPageProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "errorpageprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
