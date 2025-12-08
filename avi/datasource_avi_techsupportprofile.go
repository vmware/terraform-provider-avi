// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviTechSupportProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviTechSupportProfileRead,
		Schema: map[string]*schema.Schema{
			"archive_rules": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceArchiveRulesSchema(),
			},
			"event_params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTechSupportEventParamsSchema(),
			},
			"file_size_threshold": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"max_disk_size_percent": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_free_disk_required": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"no_of_techsupport_retentions": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"simultaneous_invocations": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"task_timeout": {
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
