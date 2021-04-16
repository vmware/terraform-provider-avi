/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviNsxtSegmentRuntime() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviNsxtSegmentRuntimeRead,
		Schema: map[string]*schema.Schema{
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dhcp6_ranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"dhcp_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dhcp_ranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nw_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"nw_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"opaque_network_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"segment_gw": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"segment_gw6": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"segment_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"segname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"subnet6": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier1_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vlan_ids": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"vrf_context_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
