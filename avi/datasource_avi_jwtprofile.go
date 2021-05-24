// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviJWTProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviJWTProfileRead,
		Schema: map[string]*schema.Schema{
			"is_federated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"jwks_keys": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceJWSKeySchema(),
			},
			"jwt_auth_type": {
				Type:     schema.TypeString,
				Computed: true,
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
