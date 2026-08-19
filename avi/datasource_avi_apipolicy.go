// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviApiPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApiPolicyRead,
		Schema: map[string]*schema.Schema{
			"active_api_labels": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiLabelsSchema(),
			},
			"api_spec_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiSpecInfoSchema(),
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
			"file_object_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"label_mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceApiPolicyLabelActionMappingSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"non_api_url_labels": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiLabelsSchema(),
			},
			"orphan_api_classification_settings": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceOrphanApiClassificationSettingsSchema(),
			},
			"orphan_api_labels": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiLabelsSchema(),
			},
			"path_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"routing_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiRoutingInfoSchema(),
			},
			"server_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiServerInfoSchema(),
			},
			"shadow_api_labels": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiLabelsSchema(),
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
			"validation_settings": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiValidationSettingsSchema(),
			},
			"zombie_api_classification_settings": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceZombieApiClassificationSettingsSchema(),
			},
			"zombie_api_labels": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiLabelsSchema(),
			},
		},
	}
}
