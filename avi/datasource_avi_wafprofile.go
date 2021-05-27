// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviWafProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafProfileRead,
		Schema: map[string]*schema.Schema{
			"config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWafConfigSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"files": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafDataFileSchema(),
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
