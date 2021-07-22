// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviServer() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServerRead,
		Schema: map[string]*schema.Schema{
			"pool_ref": {
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": {
				Type:     schema.TypeString,
				Required: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"autoscaling_group_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"availability_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"discovered_networks": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_orchestration_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hostname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"location": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"mac_address": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nw_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prst_hdr_val": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ratio": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resolve_server_by_dns": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"rewrite_host_header": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_node": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"static": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"verify_network": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vm_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
