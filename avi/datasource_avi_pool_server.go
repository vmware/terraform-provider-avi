/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviServer() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServerRead,
		Schema: map[string]*schema.Schema{
			"pool_ref": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ip": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "V4",
			},
			"autoscaling_group_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"discovered_networks": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDiscoveredNetworkSchema(),
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"external_orchestration_id": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"external_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hostname": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"location": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGeoLocationSchema(),
			},
			"mac_address": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nw_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"prst_hdr_val": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},
			"resolve_server_by_dns": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"rewrite_host_header": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"server_node": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"static": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"verify_network": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"vm_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
