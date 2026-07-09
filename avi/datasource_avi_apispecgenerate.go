// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviApiSpecGenerate() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApiSpecGenerateRead,
		Schema: map[string]*schema.Schema{
			"completed_events": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiSpecGenerateParamsSchema(),
			},
			"path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"progress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiSpecGenerateStateSchema(),
			},
			"task_events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceTaskEventMapSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"total_events": {
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
