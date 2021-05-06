// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviControllerPortalRegistration() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviControllerPortalRegistrationRead,
		Schema: map[string]*schema.Schema{
			"asset": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerPortalAssetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portal_auth": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerPortalAuthSchema(),
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
