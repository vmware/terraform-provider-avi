// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviUserAccountProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviUserAccountProfileRead,
		Schema: map[string]*schema.Schema{
			"complexity_constraint": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceComplexityConstraintSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"expiration_constraint": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceExpirationConstraintSchema(),
			},
			"lockout_constraint": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceLockoutConstraintSchema(),
			},
			"max_concurrent_sessions": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
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
