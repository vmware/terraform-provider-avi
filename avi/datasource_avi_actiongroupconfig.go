/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviActionGroupConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviActionGroupConfigRead,
		Schema: map[string]*schema.Schema{
			"action_script_config_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"autoscale_trigger_notification": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"email_config_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_only": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
			},
			"level": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"snmp_trap_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"syslog_config_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
