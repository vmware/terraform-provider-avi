/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviBackup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviBackupRead,
		Schema: map[string]*schema.Schema{
			"backup_config_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"local_file_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"remote_file_url": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scheduler_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timestamp": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
