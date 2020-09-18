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
			"application_persistence_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_health_status_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"domain_names": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"down_response": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceGslbServiceDownResponseSchema(),
			},
			"enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"groups": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceGslbPoolSchema(),
			},
			"health_monitor_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"health_monitor_scope": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hm_off": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_federated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceKeyValueSchema(),
			},
			"min_members": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_dns_ip": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pool_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resolve_cname": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"site_persistence_enabled": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"use_edns_client_subnet": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_match": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}
