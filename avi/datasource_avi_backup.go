// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviBackup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviBackupRead,
		Schema: map[string]*schema.Schema{
			"backup_config_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"local_file_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_file_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduler_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timestamp": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
