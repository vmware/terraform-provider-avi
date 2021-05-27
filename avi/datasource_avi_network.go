// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviNetwork() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviNetworkRead,
		Schema: map[string]*schema.Schema{
			"attrs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceKeyValueSchema(),
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configured_subnets": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceSubnetSchema(),
			},
			"dhcp_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"exclude_discovered_subnets": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ip6_autocfg_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
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
			"synced_from_se": {
				Type:     schema.TypeBool,
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
			"vcenter_dvs": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"vrf_context_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
