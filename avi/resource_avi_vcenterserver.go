// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceVCenterServerSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_ref": {
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
		"content_lib": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceContentLibConfigSchema(),
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
		"vcenter_credentials_ref": {
			Type:     schema.TypeString,
			Required: true,
		},
		"vcenter_url": {
			Type:     schema.TypeString,
			Required: true,
		},
	}
}

func resourceAviVCenterServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviVCenterServerCreate,
		Read:   ResourceAviVCenterServerRead,
		Update: resourceAviVCenterServerUpdate,
		Delete: resourceAviVCenterServerDelete,
		Schema: ResourceVCenterServerSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceVCenterServerImporter,
		},
	}
}

func ResourceVCenterServerImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceVCenterServerSchema()
	return ResourceImporter(d, m, "vcenterserver", s)
}

func ResourceAviVCenterServerRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVCenterServerSchema()
	err := APIRead(d, meta, "vcenterserver", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviVCenterServerCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVCenterServerSchema()
	err := APICreate(d, meta, "vcenterserver", s)
	if err == nil {
		err = ResourceAviVCenterServerRead(d, meta)
	}
	return err
}

func resourceAviVCenterServerUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceVCenterServerSchema()
	var err error
	err = APIUpdate(d, meta, "vcenterserver", s)
	if err == nil {
		err = ResourceAviVCenterServerRead(d, meta)
	}
	return err
}

func resourceAviVCenterServerDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "vcenterserver")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
