// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviApplicationInsightsPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApplicationInsightsPolicyRead,
		Schema: map[string]*schema.Schema{
			"application_insights_params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApplicationInsightsParamsSchema(),
			},
			"application_sampling_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApplicationSamplingConfigSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_application_insights": {
				Type:     schema.TypeString,
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
