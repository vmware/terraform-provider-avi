/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceAviPingAccessAgent() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPingAccessAgentRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pingaccess_pool_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"primary_server": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePoolServerSchema(),
			},
			"properties_file_data": {
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
		},
	}
}
