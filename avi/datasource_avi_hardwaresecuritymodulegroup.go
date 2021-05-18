// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviHardwareSecurityModuleGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHardwareSecurityModuleGroupRead,
		Schema: map[string]*schema.Schema{
			"hsm": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHardwareSecurityModuleSchema(),
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
