/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviTenant() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviTenantRead,
		Schema: map[string]*schema.Schema{
			"config_settings": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTenantConfigurationSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enforce_label_group": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"label_group_refs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"local": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
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
