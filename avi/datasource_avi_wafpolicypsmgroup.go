/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceAviWafPolicyPSMGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafPolicyPSMGroupRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"hit_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_learning_group": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"locations": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafPSMLocationSchema(),
			},
			"miss_action": {
				Type:     schema.TypeString,
				Computed: true,
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
