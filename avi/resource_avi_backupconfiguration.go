// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceBackupConfigurationSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"aws_access_key": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"aws_bucket_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"aws_bucket_region": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"aws_secret_access": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"backup_file_prefix": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"backup_passphrase": {
			Type:             schema.TypeString,
			Optional:         true,
			Computed:         true,
			Sensitive:        true,
			DiffSuppressFunc: suppressSensitiveFieldDiffs,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"maximum_backups_stored": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "4",
			ValidateFunc: validateInteger,
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
		"remote_file_transfer_protocol": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "SCP",
		},
		"remote_hostname": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"s3_bucket_folder": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"save_local": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
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
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"upload_to_s3": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
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
	err := APIRead(d, meta, "backupconfiguration", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviBackupConfigurationCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupConfigurationSchema()
	err := APICreate(d, meta, "backupconfiguration", s)
	if err == nil {
		err = ResourceAviBackupConfigurationRead(d, meta)
	}
	return err
}

func resourceAviBackupConfigurationUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceBackupConfigurationSchema()
	var err error
	err = APIUpdate(d, meta, "backupconfiguration", s)
	if err == nil {
		err = ResourceAviBackupConfigurationRead(d, meta)
	}
	return err
}

func resourceAviBackupConfigurationDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "backupconfiguration")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
