/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

//nolint
func dataSourceAviIpAddrGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviIpAddrGroupRead,
		Schema: map[string]*schema.Schema{
			"addrs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIpAddrSchema(),
			},
			"apic_epg_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"country_codes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_ports": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIpAddrPortSchema(),
			},
			"marathon_app_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"marathon_service_port": {
				Type:     schema.TypeInt,
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
			"prefixes": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIpAddrPrefixSchema(),
			},
			"ranges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIpAddrRangeSchema(),
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
