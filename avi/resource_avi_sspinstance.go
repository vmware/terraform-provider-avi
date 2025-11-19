// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSspInstanceSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"avi_client_cert": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"client_cert": {
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
		"feature": {
			Type:     schema.TypeString,
			Required: true,
		},
		"hostname": {
			Type:     schema.TypeString,
			Required: true,
		},
		"ingress_cert": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"resources": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSspResourcesSchema(),
		},
		"status": {
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

func resourceAviSspInstance() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSspInstanceCreate,
		Read:   ResourceAviSspInstanceRead,
		Update: resourceAviSspInstanceUpdate,
		Delete: resourceAviSspInstanceDelete,
		Schema: ResourceSspInstanceSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSspInstanceImporter,
		},
	}
}

func ResourceSspInstanceImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSspInstanceSchema()
	return ResourceImporter(d, m, "sspinstance", s)
}

func ResourceAviSspInstanceRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSspInstanceSchema()
	err := APIRead(d, meta, "sspinstance", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSspInstanceCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSspInstanceSchema()
	err := APICreate(d, meta, "sspinstance", s)
	if err == nil {
		err = ResourceAviSspInstanceRead(d, meta)
	}
	return err
}

func resourceAviSspInstanceUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSspInstanceSchema()
	var err error
	err = APIUpdate(d, meta, "sspinstance", s)
	if err == nil {
		err = ResourceAviSspInstanceRead(d, meta)
	}
	return err
}

func resourceAviSspInstanceDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "sspinstance")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
