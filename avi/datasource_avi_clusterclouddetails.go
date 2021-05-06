// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviClusterCloudDetails() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviClusterCloudDetailsRead,
		Schema: map[string]*schema.Schema{
			"azure_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAzureClusterInfoSchema(),
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
