// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviVSDataScriptSet() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVSDataScriptSetRead,
		Schema: map[string]*schema.Schema{
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"datascript": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceVSDataScriptSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_reputation_db_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ipgroup_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pool_group_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"pool_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"protocol_parser_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"rate_limiters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRateLimiterSchema(),
			},
			"string_group_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
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
