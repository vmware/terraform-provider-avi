// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviRetentionPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviRetentionPolicyRead,
		Schema: map[string]*schema.Schema{
			"enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"history": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRetentionSummarySchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePolicySpecSchema(),
			},
			"summary": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRetentionSummarySchema(),
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
