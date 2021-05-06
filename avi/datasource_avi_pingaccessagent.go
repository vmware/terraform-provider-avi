// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviPingAccessAgent() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPingAccessAgentRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pingaccess_pool_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_server": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePoolServerSchema(),
			},
			"properties_file_data": {
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
		},
	}
}
