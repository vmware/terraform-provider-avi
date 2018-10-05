/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviWafProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafProfileRead,
		Schema: map[string]*schema.Schema{
			"config": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceWafConfigSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"files": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafDataFileSchema(),
			},
			"name": &schema.Schema{
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
			},
		},
	}
}
