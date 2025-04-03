// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviAvailabilityZone() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAvailabilityZoneRead,
		Schema: map[string]*schema.Schema{
			"az_clusters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceAZClusterSchema(),
			},
			"az_datastore": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAZDatastoreSchema(),
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
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
