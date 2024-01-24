// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

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
			"client_behavior_detector": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceBotConfigClientBehaviorSchema(),
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
			"system_bot_mapping_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"system_consolidator_ref": {
				Type:     schema.TypeString,
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
			"user_bot_mapping_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_consolidator_ref": {
				Type:     schema.TypeString,
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
