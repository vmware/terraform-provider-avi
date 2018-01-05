/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCloudProperties() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCloudPropertiesRead,
		Schema: map[string]*schema.Schema{
			"cc_props": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceCC_PropertiesSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"cc_vtypes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"hyp_props": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceHypervisor_PropertiesSchema(),
			},
			"info": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloudInfoSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
