/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviScheduler() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSchedulerRead,
		Schema: map[string]*schema.Schema{
			"backup_config_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"end_date_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"frequency": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"frequency_unit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"run_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"run_script_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"scheduler_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_date_time": {
				Type:     schema.TypeString,
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
