// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviGeoDB() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGeoDBRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"files": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceGeoDBFileSchema(),
			},
			"is_federated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceGeoDBMappingSchema(),
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
