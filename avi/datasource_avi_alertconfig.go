/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviAlertConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAlertConfigRead,
		Schema: map[string]*schema.Schema{
			"action_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alert_rule": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAlertRuleSchema(),
			},
			"autoscale_alert": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"expiry_time": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"obj_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"object_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"recommendation": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rolling_window": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"source": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"summary": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"threshold": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"throttle": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  600,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
