/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviALBServicesConfig() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviALBServicesConfigRead,
		Schema: map[string]*schema.Schema{
			"asset_contact": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceALBServicesUserSchema(),
			},
			"feature_opt_in_status": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePortalFeatureOptInSchema(),
			},
			"ip_reputation_file_object_expiry_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ip_reputation_sync_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"polling_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"portal_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"proactive_support_defaults": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProactiveSupportDefaultsSchema(),
			},
			"split_proxy_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProxyConfigurationSchema(),
			},
			"use_system_proxy": {
				Type:     schema.TypeBool,
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
