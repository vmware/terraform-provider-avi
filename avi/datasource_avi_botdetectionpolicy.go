/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviBotDetectionPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviBotDetectionPolicyRead,
		Schema: map[string]*schema.Schema{
			"allow_list": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceBotAllowListSchema(),
			},
			"bot_mapping_uuids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"consolidator_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_location_detector": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceBotConfigIPLocationSchema(),
			},
			"ip_reputation_detector": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceBotConfigIPReputationSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_agent_detector": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceBotConfigUserAgentSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
