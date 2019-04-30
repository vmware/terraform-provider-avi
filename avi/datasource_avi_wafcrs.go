/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviWafCRS() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafCRSRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafRuleGroupSchema(),
			},
			"integrity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"release_date": &schema.Schema{
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
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
