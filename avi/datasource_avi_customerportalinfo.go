// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviCustomerPortalInfo() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCustomerPortalInfoRead,
		Schema: map[string]*schema.Schema{
			"polling_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"portal_url": {
				Type:     schema.TypeString,
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
