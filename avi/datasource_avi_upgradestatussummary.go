// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviUpgradeStatusSummary() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviUpgradeStatusSummaryRead,
		Schema: map[string]*schema.Schema{
			"enable_patch_rollback": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_rollback": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"obj_cloud_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_image_ref": {
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
				Elem:     ResourceUpgradeOpsStateSchema(),
			},
			"tasks_completed": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"total_tasks": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upgrade_ops": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
