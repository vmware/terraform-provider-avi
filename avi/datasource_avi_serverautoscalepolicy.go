/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviServerAutoScalePolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServerAutoScalePolicyRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"intelligent_autoscale": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"intelligent_scalein_margin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  40,
			},
			"intelligent_scaleout_margin": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"max_scalein_adjustment_step": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"max_scaleout_adjustment_step": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"max_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"min_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"scalein_alertconfig_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"scalein_cooldown": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"scaleout_alertconfig_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"scaleout_cooldown": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"use_predicted_load": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
