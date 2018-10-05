/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviGslb() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbRead,
		Schema: map[string]*schema.Schema{
			"clear_on_max_retries": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  20,
			},
			"client_ip_addr_group": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbClientIpAddrGroupSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_configs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceDNSConfigSchema(),
			},
			"is_federated": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"leader_cluster_uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"maintenance_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"send_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  15,
			},
			"sites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbSiteSchema(),
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"third_party_sites": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbThirdPartySiteSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"view_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
		},
	}
}
