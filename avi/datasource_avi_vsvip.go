/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviVsVip() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVsVipRead,
		Schema: map[string]*schema.Schema{
			"bgp_peer_labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"dns_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDnsInfoSchema(),
			},
			"east_west_placement": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ipam_selector": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSelectorSchema(),
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
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tier1_lr": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_standard_alb": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vip": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceVipSchema(),
			},
			"vrf_context_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vsvip_cloud_config_cksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
