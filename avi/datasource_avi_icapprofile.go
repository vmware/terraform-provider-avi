/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviIcapProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviIcapProfileRead,
		Schema: map[string]*schema.Schema{
			"buffer_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"buffer_size_exceed_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_preview": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"fail_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_group_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"preview_size": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"response_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"service_uri": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"slow_response_warning_threshold": {
				Type:     schema.TypeInt,
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
			"vendor": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
