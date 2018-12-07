/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCluster() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviClusterRead,
		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"nodes": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceClusterNodeSchema(),
			},
			"rejoin_nodes_automatically": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"virtual_ip": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpAddrSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
		},
	}
}
