/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviAutoScaleLaunchConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAutoScaleLaunchConfigRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"image_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mesos": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAutoScaleMesosSettingsSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"openstack": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAutoScaleOpenStackSettingsSchema(),
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_external_asg": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
