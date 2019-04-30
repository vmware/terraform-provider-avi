/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPingAccessAgent() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPingAccessAgentRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pingaccess_pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"primary_server": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcePoolServerSchema(),
			},
			"properties_file_data": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
