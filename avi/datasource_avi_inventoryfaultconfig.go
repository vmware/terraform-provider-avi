// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviInventoryFaultConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviInventoryFaultConfigRead,
		Schema: map[string]*schema.Schema{
			"controller_faults": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerFaultsSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serviceengine_faults": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceServiceengineFaultsSchema(),
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
			"virtualservice_faults": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceVirtualserviceFaultsSchema(),
			},
		},
	}
}
