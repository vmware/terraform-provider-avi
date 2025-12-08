// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviUpgradeProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviUpgradeProfileRead,
		Schema: map[string]*schema.Schema{
			"controller_params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerParamsSchema(),
			},
			"dry_run": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDryRunParamsSchema(),
			},
			"image": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceImageParamsSchema(),
			},
			"pre_checks": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePreChecksParamsSchema(),
			},
			"service_engine": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceServiceEngineParamsSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
