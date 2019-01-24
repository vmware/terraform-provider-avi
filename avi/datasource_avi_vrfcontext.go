/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviVrfContext() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVrfContextRead,
		Schema: map[string]*schema.Schema{
			"bgp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceBgpProfileSchema(),
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"debugvrfcontext": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDebugVrfContextSchema(),
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"gateway_mon": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGatewayMonitorSchema(),
			},
			"internal_gateway_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceInternalGatewayMonitorSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"static_routes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceStaticRouteSchema(),
			},
			"system_default": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
