// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviGslbService() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbServiceRead,
		Schema: map[string]*schema.Schema{
			"application_persistence_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_health_status_enabled": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_federated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"min_members": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"num_dns_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pool_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"resolve_cname": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_persistence_enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ttl": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"use_edns_client_subnet": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wildcard_match": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
