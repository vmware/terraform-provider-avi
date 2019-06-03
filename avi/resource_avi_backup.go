/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform/helper/schema"
	"log"
	"strings"
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
		},
		"remote_file_url": {
			Type:     schema.TypeString,
			Optional: true,
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
	err := ApiRead(d, meta, "backup", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviBackupCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupSchema()
	err := ApiCreateOrUpdate(d, meta, "backup", s)
	if err == nil {
		err = ResourceAviBackupRead(d, meta)
	}
	return err
}

func resourceAviBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "backup", s)
	if err == nil {
		err = ResourceAviBackupRead(d, meta)
	}
	return err
}

func resourceAviBackupDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "backup"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
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
