// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviWafApplicationSignatureProvider() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafApplicationSignatureProviderRead,
		Schema: map[string]*schema.Schema{
			"available_applications": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafApplicationSignatureAppVersionSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"filter_rules_on_import": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ruleset_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_status": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceWebApplicationSignatureServiceStatusSchema(),
			},
			"signatures": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceWafRuleSchema(),
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
