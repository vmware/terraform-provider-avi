/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviScheduler() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSchedulerRead,
		Schema: map[string]*schema.Schema{
			"backup_config_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"end_date_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"frequency": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"frequency_unit": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"run_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"run_script_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scheduler_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SCHEDULER_ACTION_BACKUP"},
			"start_date_time": &schema.Schema{
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
			},
		},
	}
}
