/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviApplicationPersistenceProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApplicationPersistenceProfileRead,
		Schema: map[string]*schema.Schema{
			"app_cookie_persistence_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAppCookiePersistenceProfileSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hdr_persistence_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHdrPersistenceProfileSchema(),
			},
			"http_cookie_persistence_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHttpCookiePersistenceProfileSchema(),
			},
			"ip_persistence_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIPPersistenceProfileSchema(),
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
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"persistence_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_hm_down_recovery": {
				Type:     schema.TypeString,
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
		},
	}
}
