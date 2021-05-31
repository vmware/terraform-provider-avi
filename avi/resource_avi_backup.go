// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/vmware/alb-sdk/go/clients"
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
	err := APICreateOrUpdate(d, meta, "backup", s)
	if err == nil {
		err = ResourceAviBackupRead(d, meta)
	}
	return err
}

func resourceAviBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "backup", s)
	if err == nil {
		err = ResourceAviBackupRead(d, meta)
	}
	return err
}

func resourceAviBackupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "backup"
	client := meta.(*clients.AviClient)
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviBackupDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
