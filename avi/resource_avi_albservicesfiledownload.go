// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceALBServicesFileDownloadSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"destination_dir": {
			Type:     schema.TypeString,
			Required: true,
		},
		"file_type": {
			Type:     schema.TypeString,
			Required: true,
		},
		"file_uri": {
			Type:     schema.TypeString,
			Required: true,
		},
		"message": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"metadata": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceALBServicesFileDownloadMetadataSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
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

func resourceAviALBServicesFileDownload() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviALBServicesFileDownloadCreate,
		Read:   ResourceAviALBServicesFileDownloadRead,
		Update: resourceAviALBServicesFileDownloadUpdate,
		Delete: resourceAviALBServicesFileDownloadDelete,
		Schema: ResourceALBServicesFileDownloadSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceALBServicesFileDownloadImporter,
		},
	}
}

func ResourceALBServicesFileDownloadImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceALBServicesFileDownloadSchema()
	return ResourceImporter(d, m, "albservicesfiledownload", s)
}

func ResourceAviALBServicesFileDownloadRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesFileDownloadSchema()
	err := APIRead(d, meta, "albservicesfiledownload", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviALBServicesFileDownloadCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesFileDownloadSchema()
	err := APICreateOrUpdate(d, meta, "albservicesfiledownload", s)
	if err == nil {
		err = ResourceAviALBServicesFileDownloadRead(d, meta)
	}
	return err
}

func resourceAviALBServicesFileDownloadUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceALBServicesFileDownloadSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "albservicesfiledownload", s)
	if err == nil {
		err = ResourceAviALBServicesFileDownloadRead(d, meta)
	}
	return err
}

func resourceAviALBServicesFileDownloadDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "albservicesfiledownload")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
