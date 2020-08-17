/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviNetwork() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviNetworkRead,
		Schema: map[string]*schema.Schema{
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
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceKeyValueSchema(),
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
