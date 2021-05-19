/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviActionGroupConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviActionGroupConfigRead,
		Schema: map[string]*schema.Schema{
			"action_script_config_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"autoscale_trigger_notification": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_config_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_only": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"level": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"snmp_trap_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"syslog_config_ref": {
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
