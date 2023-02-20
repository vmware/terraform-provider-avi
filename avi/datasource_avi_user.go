// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviUser() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviUserRead,
		Schema: map[string]*schema.Schema{
			"access": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceUserRoleSchema(),
			},
			"anonymous_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"date_joined": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"default_tenant_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"email": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"full_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_active": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_internal_user": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_staff": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_superuser": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"local": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"logged_in": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"passwordless": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"recovery_token": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"token_expiration_date": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ui_property": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unix_crypt_password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"user_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"username": {
				Type:     schema.TypeString,
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
