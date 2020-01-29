/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import (
	"github.com/avinetworks/sdk/go/clients"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"log"
	"strings"
)

func ResourceBackupConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"aws_access_key": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"aws_bucket_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"aws_secret_access": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"backup_file_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"backup_passphrase": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"maximum_backups_stored": {
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"remote_directory": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"remote_hostname": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"save_local": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"ssh_user_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"upload_to_remote_host": {
			Type:     schema.TypeBool,
			Optional: true,
			Computed: true,
		},
		"upload_to_s3": {
			Type:     schema.TypeBool,
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

func resourceAviBackupConfiguration() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviBackupConfigurationCreate,
		Read:   ResourceAviBackupConfigurationRead,
		Update: resourceAviBackupConfigurationUpdate,
		Delete: resourceAviBackupConfigurationDelete,
		Schema: ResourceBackupConfigurationSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceBackupConfigurationImporter,
		},
	}
}

func ResourceBackupConfigurationImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceBackupConfigurationSchema()
	return ResourceImporter(d, m, "backupconfiguration", s)
}

func ResourceAviBackupConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupConfigurationSchema()
	err := ApiRead(d, meta, "backupconfiguration", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviBackupConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupConfigurationSchema()
	err := ApiCreateOrUpdate(d, meta, "backupconfiguration", s)
	if err == nil {
		err = ResourceAviBackupConfigurationRead(d, meta)
	}
	return err
}

func resourceAviBackupConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupConfigurationSchema()
	var err error
	err = ApiCreateOrUpdate(d, meta, "backupconfiguration", s)
	if err == nil {
		err = ResourceAviBackupConfigurationRead(d, meta)
	}
	return err
}

func resourceAviBackupConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "backupconfiguration"
	if ApiDeleteSystemDefaultCheck(d) {
		return nil
	}
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204") || strings.Contains(err.Error(), "403")) {
			log.Println("[INFO] resourceAviBackupConfigurationDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
