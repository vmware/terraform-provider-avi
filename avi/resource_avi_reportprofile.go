// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceReportProfileSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"collection_rules": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceCollectionRulesSchema(),
		},
		"max_concurrent_reports": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "1",
			ValidateFunc: validateInteger,
		},
		"remote_controller": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceRemoteControllerSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviReportProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviReportProfileCreate,
		Read:   ResourceAviReportProfileRead,
		Update: resourceAviReportProfileUpdate,
		Delete: resourceAviReportProfileDelete,
		Schema: ResourceReportProfileSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceReportProfileImporter,
		},
	}
}

func ResourceReportProfileImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceReportProfileSchema()
	return ResourceImporter(d, m, "reportprofile", s)
}

func ResourceAviReportProfileRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceReportProfileSchema()
	err := APIRead(d, meta, "reportprofile", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviReportProfileCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceReportProfileSchema()
	err := APICreate(d, meta, "reportprofile", s)
	if err == nil {
		err = ResourceAviReportProfileRead(d, meta)
	}
	return err
}

func resourceAviReportProfileUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceReportProfileSchema()
	var err error
	err = APIUpdate(d, meta, "reportprofile", s)
	if err == nil {
		err = ResourceAviReportProfileRead(d, meta)
	}
	return err
}

func resourceAviReportProfileDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "reportprofile")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
