/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviServerAutoScalePolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServerAutoScalePolicyRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"delay_for_server_garbage_collection": {
				Type:     schema.TypeInt,
				Computed: true,
			},
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
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
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
			"scheduled_scalings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceScheduledScalingSchema(),
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
