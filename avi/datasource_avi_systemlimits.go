/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSystemLimits() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSystemLimitsRead,
		Schema: map[string]*schema.Schema{
			"controller_limits": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerLimitsSchema(),
			},
			"controller_sizes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceControllerSizeSchema(),
			},
			"serviceengine_limits": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceServiceEngineLimitsSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
