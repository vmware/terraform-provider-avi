// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceVsGsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"geodb_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"gs_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"gslb_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
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
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vs_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviVsGs() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVsGsCreate,
		Read:   ResourceAviVsGsRead,
		Update: resourceAviVsGsUpdate,
		Delete: resourceAviVsGsDelete,
		Schema: ResourceVsGsSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVsGsImporter,
		},
	}
}

func ResourceVsGsImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVsGsSchema()
	return ResourceImporter(d, m, "vsgs", s)
}

func ResourceAviVsGsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsGsSchema()
	err := APIRead(d, meta, "vsgs", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVsGsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsGsSchema()
	err := APICreateOrUpdate(d, meta, "vsgs", s)
	if err == nil {
		err = ResourceAviVsGsRead(d, meta)
	}
	return err
}

func resourceAviVsGsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVsGsSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "vsgs", s)
	if err == nil {
		err = ResourceAviVsGsRead(d, meta)
	}
	return err
}

func resourceAviVsGsDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "vsgs")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
