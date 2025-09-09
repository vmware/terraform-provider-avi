// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviReport() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviReportRead,
		Schema: map[string]*schema.Schema{
			"duration": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"filename": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pre_check": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceReadinessCheckObjSchema(),
			},
			"progress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"request": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceReportGenerationRequestSchema(),
			},
			"start_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceReportGenStateSchema(),
			},
			"tasks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceTaskEventMapSchema(),
			},
			"tasks_completed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"total_tasks": {
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
