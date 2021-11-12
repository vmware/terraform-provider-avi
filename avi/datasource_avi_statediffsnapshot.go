// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviStatediffSnapshot() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviStatediffSnapshotRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"gslb_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"gslb_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pool_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"post_snapshot": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcepostsnapshotSchema(),
			},
			"pre_snapshot": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcepresnapshotSchema(),
			},
			"se_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_group_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"snapshot_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"statediff_operation_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vs_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vs_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
