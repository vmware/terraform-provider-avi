// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSeProperties() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSePropertiesRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"se_agent_properties": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSeAgentPropertiesSchema(),
			},
			"se_bootup_properties": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSeBootupPropertiesSchema(),
			},
			"se_runtime_properties": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSeRuntimePropertiesSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
