// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviJWTServerProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviJWTServerProfileRead,
		Schema: map[string]*schema.Schema{
			"is_federated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"issuer": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"jwks_keys": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"jwt_profile_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"jwt_server_profile_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceJWTServerProfileConfigSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
