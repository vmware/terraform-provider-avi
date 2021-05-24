// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSiteVersion() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSiteVersionRead,
		Schema: map[string]*schema.Schema{
			"datetime": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prev_target_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"replication_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_timeline": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timeline": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"version_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
