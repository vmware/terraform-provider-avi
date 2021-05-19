/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviAutoScaleLaunchConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAutoScaleLaunchConfigRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"mesos": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAutoScaleMesosSettingsSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"openstack": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAutoScaleOpenStackSettingsSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_external_asg": {
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
