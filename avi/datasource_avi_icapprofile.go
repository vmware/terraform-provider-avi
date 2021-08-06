// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviIcapProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviIcapProfileRead,
		Schema: map[string]*schema.Schema{
			"allow_204": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"buffer_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"buffer_size_exceed_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_preview": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"fail_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_group_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"preview_size": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"response_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"slow_response_warning_threshold": {
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
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
