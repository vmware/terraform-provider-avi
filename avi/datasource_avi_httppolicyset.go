// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviHTTPPolicySet() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHTTPPolicySetRead,
		Schema: map[string]*schema.Schema{
			"cloud_config_cksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_request_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTPRequestPolicySchema(),
			},
			"http_response_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTPResponsePolicySchema(),
			},
			"http_security_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTPSecurityPolicySchema(),
			},
			"is_internal_policy": {
				Type:     schema.TypeBool,
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
