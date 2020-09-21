/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviRole() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviRoleRead,
		Schema: map[string]*schema.Schema{
			"filters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privileges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourcePermissionSchema(),
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
