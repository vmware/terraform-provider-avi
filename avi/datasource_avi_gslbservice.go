/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviGslbService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbServiceRead,
		Schema: map[string]*schema.Schema{
			"application_persistence_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"controller_health_status_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"domain_names": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"down_response": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGslbServiceDownResponseSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceGslbPoolSchema(),
			},
			"health_monitor_refs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"health_monitor_scope": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_SERVICE_HEALTH_MONITOR_ALL_MEMBERS"},
			"is_federated": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"min_members": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"num_dns_ip": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"pool_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "GSLB_SERVICE_ALGORITHM_PRIORITY"},
			"site_persistence_enabled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ttl": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"use_edns_client_subnet": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"wildcard_match": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}
