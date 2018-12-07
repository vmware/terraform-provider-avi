/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviBackupConfiguration() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviBackupConfigurationRead,
		Schema: map[string]*schema.Schema{
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
				Optional: true,
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
				Computed: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_to_remote_host": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
