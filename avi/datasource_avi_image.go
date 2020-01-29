/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviImage() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviImageRead,
		Schema: map[string]*schema.Schema{
			"controller_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePackageDetailsSchema(),
			},
			"migrations": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSupportedMigrationsSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"se_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePackageDetailsSchema(),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
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
