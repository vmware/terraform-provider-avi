// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviReportProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviReportProfileRead,
		Schema: map[string]*schema.Schema{
			"collection_rules": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceCollectionRulesSchema(),
			},
			"max_concurrent_reports": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"remote_controller": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRemoteControllerSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
