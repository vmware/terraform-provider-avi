// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviCloudProperties() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCloudPropertiesRead,
		Schema: map[string]*schema.Schema{
			"cc_props": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceCC_PropertiesSchema(),
			},
			"cc_vtypes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hyp_props": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceHypervisor_PropertiesSchema(),
			},
			"info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCloudInfoSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
