// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceAvailabilityZoneSchema() map[string]*schema.Schema {
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
		"vcenter_refs": {
			Type:     schema.TypeList,
			Required: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func resourceAviAvailabilityZone() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviAvailabilityZoneCreate,
		Read:   ResourceAviAvailabilityZoneRead,
		Update: resourceAviAvailabilityZoneUpdate,
		Delete: resourceAviAvailabilityZoneDelete,
		Schema: ResourceAvailabilityZoneSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceAvailabilityZoneImporter,
		},
	}
}

func ResourceAvailabilityZoneImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceAvailabilityZoneSchema()
	return ResourceImporter(d, m, "availabilityzone", s)
}

func ResourceAviAvailabilityZoneRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAvailabilityZoneSchema()
	err := APIRead(d, meta, "availabilityzone", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviAvailabilityZoneCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAvailabilityZoneSchema()
	err := APICreate(d, meta, "availabilityzone", s)
	if err == nil {
		err = ResourceAviAvailabilityZoneRead(d, meta)
	}
	return err
}

func resourceAviAvailabilityZoneUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceAvailabilityZoneSchema()
	var err error
	err = APIUpdate(d, meta, "availabilityzone", s)
	if err == nil {
		err = ResourceAviAvailabilityZoneRead(d, meta)
	}
	return err
}

func resourceAviAvailabilityZoneDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "availabilityzone")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
