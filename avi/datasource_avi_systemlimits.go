// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSystemLimits() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSystemLimitsRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"controller_limits": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerLimitsSchema(),
			},
			"controller_sizes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceControllerSizeSchema(),
			},
			"serviceengine_limits": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceServiceEngineLimitsSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
