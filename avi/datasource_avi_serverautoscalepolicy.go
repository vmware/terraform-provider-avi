/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceAviServerAutoScalePolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServerAutoScalePolicyRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"intelligent_autoscale": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"intelligent_scalein_margin": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"intelligent_scaleout_margin": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_scalein_adjustment_step": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_scaleout_adjustment_step": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"min_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scalein_alertconfig_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"scalein_cooldown": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"scaleout_alertconfig_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"scaleout_cooldown": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_predicted_load": {
				Type:     schema.TypeBool,
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
