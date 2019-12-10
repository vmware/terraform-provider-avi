/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

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
