// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceMicroServiceGroupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"created_by": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"service_refs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
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

func resourceAviMicroServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviMicroServiceGroupCreate,
		Read:   ResourceAviMicroServiceGroupRead,
		Update: resourceAviMicroServiceGroupUpdate,
		Delete: resourceAviMicroServiceGroupDelete,
		Schema: ResourceMicroServiceGroupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceMicroServiceGroupImporter,
		},
	}
}

func ResourceMicroServiceGroupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceMicroServiceGroupSchema()
	return ResourceImporter(d, m, "microservicegroup", s)
}

func ResourceAviMicroServiceGroupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceMicroServiceGroupSchema()
	err := APIRead(d, meta, "microservicegroup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviMicroServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceMicroServiceGroupSchema()
	err := APICreateOrUpdate(d, meta, "microservicegroup", s)
	if err == nil {
		err = ResourceAviMicroServiceGroupRead(d, meta)
	}
	return err
}

func resourceAviMicroServiceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceMicroServiceGroupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "microservicegroup", s)
	if err == nil {
		err = ResourceAviMicroServiceGroupRead(d, meta)
	}
	return err
}

func resourceAviMicroServiceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "microservicegroup")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
