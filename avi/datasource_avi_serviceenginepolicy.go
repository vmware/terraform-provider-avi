/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviServiceEnginePolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServiceEnginePolicyRead,
		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nat_policy_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_group_ref": {
				Type:     schema.TypeString,
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
			"vrf_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
