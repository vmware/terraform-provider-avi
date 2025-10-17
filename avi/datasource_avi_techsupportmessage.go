// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviTechSupportMessage() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviTechSupportMessageRead,
		Schema: map[string]*schema.Schema{
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_code": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tech_support_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
