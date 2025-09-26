// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviPasswordPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPasswordPolicyRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"lockout_evaluation_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lockout_max_auth_failures": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"lockout_period": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_length": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_lowercase": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_numeric": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_special": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_uppercase": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password_expiration_days": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"password_history": {
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
