/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviPriorityLabels() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPriorityLabelsRead,
		Schema: map[string]*schema.Schema{
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
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"equivalent_labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceEquivalentLabelsSchema(),
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
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
