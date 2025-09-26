// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviApplicationInsightsState() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApplicationInsightsStateRead,
		Schema: map[string]*schema.Schema{
			"application_insights_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"application_sampling_runtime": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApplicationSamplingRuntimeSchema(),
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
