// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviImage() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviImageRead,
		Schema: map[string]*schema.Schema{
			"cloud_info_values": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceImageCloudDataSchema(),
			},
			"controller_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePackageDetailsSchema(),
			},
			"controller_patch_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_patch_uuid": {
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
			"events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceImageEventMapSchema(),
			},
			"img_state": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceImageUploadOpsStatusSchema(),
			},
			"migrations": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSupportedMigrationsSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"progress": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePackageDetailsSchema(),
			},
			"se_patch_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_patch_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeString,
				Computed: true,
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
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uber_bundle": {
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
