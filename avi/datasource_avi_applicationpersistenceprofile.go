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
			"app_cookie_persistence_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAppCookiePersistenceProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"hdr_persistence_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHdrPersistenceProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"http_cookie_persistence_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHttpCookiePersistenceProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ip_persistence_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIPPersistenceProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"is_federated": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"persistence_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"server_hm_down_recovery": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "HM_DOWN_PICK_NEW_SERVER"},
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
