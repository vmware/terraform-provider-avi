// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviTechSupport() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviTechSupportRead,
		Schema: map[string]*schema.Schema{
			"case_number": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
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
			"errors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"level": {
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
			"obj_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"obj_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"output": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTechSupportParamsSchema(),
			},
			"progress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"size": {
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
				Elem:     ResourceTechSupportStateSchema(),
			},
			"tasks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceTechSupportEventMapSchema(),
			},
			"tasks_completed": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"techsupport_readiness": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceReadinessCheckObjSchema(),
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
			"warnings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}
