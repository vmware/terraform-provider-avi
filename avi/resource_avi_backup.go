// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceBackupSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"backup_config_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"file_name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"local_file_url": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"remote_file_url": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"scheduler_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"timestamp": {
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

func resourceAviBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviBackupCreate,
		Read:   ResourceAviBackupRead,
		Update: resourceAviBackupUpdate,
		Delete: resourceAviBackupDelete,
		Schema: ResourceBackupSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceBackupImporter,
		},
	}
}

func ResourceBackupImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceBackupSchema()
	return ResourceImporter(d, m, "backup", s)
}

func ResourceAviBackupRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupSchema()
	err := APIRead(d, meta, "backup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviBackupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupSchema()
	err := APICreate(d, meta, "backup", s)
	if err == nil {
		err = ResourceAviBackupRead(d, meta)
	}
	return err
}

func resourceAviBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupSchema()
	var err error
	err = APIUpdate(d, meta, "backup", s)
	if err == nil {
		err = ResourceAviBackupRead(d, meta)
	}
	return err
}

func resourceAviBackupDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "backup")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
