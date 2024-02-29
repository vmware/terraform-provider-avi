// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceLicenseStatusSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"saas_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSaasLicensingStatusSchema(),
		},
		"service_update": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceLicenseServiceUpdateSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviLicenseStatus() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviLicenseStatusCreate,
		Read:   ResourceAviLicenseStatusRead,
		Update: resourceAviLicenseStatusUpdate,
		Delete: resourceAviLicenseStatusDelete,
		Schema: ResourceLicenseStatusSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceLicenseStatusImporter,
		},
	}
}

func ResourceLicenseStatusImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceLicenseStatusSchema()
	return ResourceImporter(d, m, "licensestatus", s)
}

func ResourceAviLicenseStatusRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseStatusSchema()
	err := APIRead(d, meta, "licensestatus", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviLicenseStatusCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseStatusSchema()
	err := APICreate(d, meta, "licensestatus", s)
	if err == nil {
		err = ResourceAviLicenseStatusRead(d, meta)
	}
	return err
}

func resourceAviLicenseStatusUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLicenseStatusSchema()
	var err error
	err = APIUpdate(d, meta, "licensestatus", s)
	if err == nil {
		err = ResourceAviLicenseStatusRead(d, meta)
	}
	return err
}

func resourceAviLicenseStatusDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "licensestatus")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
