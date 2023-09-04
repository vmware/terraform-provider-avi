// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func ResourceALBServicesFileUploadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"case_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"error": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"file_path": {
			Type:     schema.TypeString,
			Required: true,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"s3_directory": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"status": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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

func resourceAviALBServicesFileUpload() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviALBServicesFileUploadCreate,
		Read:   ResourceAviALBServicesFileUploadRead,
		Update: resourceAviALBServicesFileUploadUpdate,
		Delete: resourceAviALBServicesFileUploadDelete,
		Schema: ResourceALBServicesFileUploadSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceALBServicesFileUploadImporter,
		},
	}
}

func ResourceALBServicesFileUploadImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceALBServicesFileUploadSchema()
	return ResourceImporter(d, m, "albservicesfileupload", s)
}

func ResourceAviALBServicesFileUploadRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesFileUploadSchema()
	err := APIRead(d, meta, "albservicesfileupload", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviALBServicesFileUploadCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesFileUploadSchema()
	err := APICreateOrUpdate(d, meta, "albservicesfileupload", s)
	if err == nil {
		err = ResourceAviALBServicesFileUploadRead(d, meta)
	}
	return err
}

func resourceAviALBServicesFileUploadUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesFileUploadSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "albservicesfileupload", s)
	if err == nil {
		err = ResourceAviALBServicesFileUploadRead(d, meta)
	}
	return err
}

func resourceAviALBServicesFileUploadDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "albservicesfileupload")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
