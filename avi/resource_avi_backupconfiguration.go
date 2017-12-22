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

func ResourceBackupConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"backup_file_prefix": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"backup_passphrase": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"maximum_backups_stored": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  4,
		},
		"name": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},
		"remote_directory": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"remote_hostname": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"save_local": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"ssh_user_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"tenant_ref": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},
		"upload_to_remote_host": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
		},
		"uuid": &schema.Schema{
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
	}
}

func ResourceAviBackupConfigurationRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupConfigurationSchema()
	client := meta.(*clients.AviClient)
	var obj interface{}
	if uuid, ok := d.GetOk("uuid"); ok {
		path := "api/backupconfiguration/" + uuid.(string)
		err := client.AviSession.Get(path, &obj)
		if err != nil {
			d.SetId("")
			return nil
		}
	} else {
		d.SetId("")
		return nil
	}
	if _, err := ApiDataToSchema(obj, d, s); err == nil {
		if err != nil {
			log.Printf("[ERROR] in setting read object %v\n", err)
		}
	}
	return nil
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
	err := ApiCreateOrUpdate(d, meta, "backupconfiguration", s)
	if err == nil {
		err = ResourceAviBackupConfigurationRead(d, meta)
	}
	return err
}

func resourceAviBackupConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	objType := "backupconfiguration"
	client := meta.(*clients.AviClient)
	uuid := d.Get("uuid").(string)
	if uuid != "" {
		path := "api/" + objType + "/" + uuid
		err := client.AviSession.Delete(path)
		if err != nil && !(strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "204")) {
			log.Println("[INFO] resourceAviBackupConfigurationDelete not found")
			return err
		}
		d.SetId("")
	}
	return nil
}
