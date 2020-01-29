/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceAviL4PolicySet() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviL4PolicySetRead,
		Schema: map[string]*schema.Schema{
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_internal_policy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"l4_connection_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceL4ConnectionPolicySchema(),
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
