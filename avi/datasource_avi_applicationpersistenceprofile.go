// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviApplicationPersistenceProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApplicationPersistenceProfileRead,
		Schema: map[string]*schema.Schema{
			"app_cookie_persistence_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAppCookiePersistenceProfileSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
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
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
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
