// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTaskJournalSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"errors": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceJournalErrorSchema(),
		},
		"image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"info": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceJournalInfoSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"obj_cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"operation": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"summary": {
			Type:     schema.TypeSet,
			Required: true,
			Elem:     ResourceJournalSummarySchema(),
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

func resourceAviTaskJournal() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTaskJournalCreate,
		Read:   ResourceAviTaskJournalRead,
		Update: resourceAviTaskJournalUpdate,
		Delete: resourceAviTaskJournalDelete,
		Schema: ResourceTaskJournalSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTaskJournalImporter,
		},
	}
}

func ResourceTaskJournalImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTaskJournalSchema()
	return ResourceImporter(d, m, "taskjournal", s)
}

func ResourceAviTaskJournalRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTaskJournalSchema()
	err := APIRead(d, meta, "taskjournal", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTaskJournalCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTaskJournalSchema()
	err := APICreateOrUpdate(d, meta, "taskjournal", s)
	if err == nil {
		err = ResourceAviTaskJournalRead(d, meta)
	}
	return err
}

func resourceAviTaskJournalUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTaskJournalSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "taskjournal", s)
	if err == nil {
		err = ResourceAviTaskJournalRead(d, meta)
	}
	return err
}

func resourceAviTaskJournalDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "taskjournal")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
