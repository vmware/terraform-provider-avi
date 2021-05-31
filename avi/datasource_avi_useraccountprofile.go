// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviUserAccountProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviUserAccountProfileRead,
		Schema: map[string]*schema.Schema{
			"account_lock_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"credentials_timeout_threshold": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_concurrent_sessions": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_login_failure_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"max_password_history_count": {
				Type:     schema.TypeInt,
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
