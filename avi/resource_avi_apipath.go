// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceApiPathSchema() map[string]*schema.Schema {
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
		"endpoints": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceApiEndpointSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"path_template": {
			Type:     schema.TypeString,
			Required: true,
			ForceNew: true,
		},
		"source": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"unknown_http_method_action": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "API_ACTION_INHERIT_FROM_API_POLICY",
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviApiPath() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApiPathCreate,
		Read:   ResourceAviApiPathRead,
		Update: resourceAviApiPathUpdate,
		Delete: resourceAviApiPathDelete,
		Schema: ResourceApiPathSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApiPathImporter,
		},
	}
}

func ResourceApiPathImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApiPathSchema()
	return ResourceImporter(d, m, "apipath", s)
}

func ResourceAviApiPathRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiPathSchema()
	err := APIRead(d, meta, "apipath", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApiPathCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiPathSchema()
	err := APICreate(d, meta, "apipath", s)
	if err == nil {
		err = ResourceAviApiPathRead(d, meta)
	}
	return err
}

func resourceAviApiPathUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiPathSchema()
	var err error
	err = APIUpdate(d, meta, "apipath", s)
	if err == nil {
		err = ResourceAviApiPathRead(d, meta)
	}
	return err
}

func resourceAviApiPathDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "apipath")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
