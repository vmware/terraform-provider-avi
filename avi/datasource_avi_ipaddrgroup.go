/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

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
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceKeyValueSchema(),
			},
			"marathon_app_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"marathon_service_port": {
				Type:     schema.TypeInt,
				Computed: true,
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
