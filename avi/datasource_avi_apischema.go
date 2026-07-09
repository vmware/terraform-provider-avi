// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviApiSchema() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApiSchemaRead,
		Schema: map[string]*schema.Schema{
			"additional_object_key_action": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"additional_properties_schema": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiSimpleSchemaDescriptionSchema(),
			},
			"allow_additional_properties": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"array_item_type": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceApiSimpleSchemaDescriptionSchema(),
			},
			"composite_types": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceApiSimpleSchemaDescriptionSchema(),
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
			"discriminator": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDiscriminatorDescriptionSchema(),
			},
			"max_items": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"min_items": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"object_properties": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceApiObjectPropertiesSchema(),
			},
			"source": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"unique_items": {
				Type:     schema.TypeString,
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
