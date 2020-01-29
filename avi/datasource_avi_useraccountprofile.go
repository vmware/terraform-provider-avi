/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceAviUserAccountProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviUserAccountProfileRead,
		Schema: map[string]*schema.Schema{
			"account_lock_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
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
