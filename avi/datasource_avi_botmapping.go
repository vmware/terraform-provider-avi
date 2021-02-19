/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviBotMapping() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviBotMappingRead,
		Schema: map[string]*schema.Schema{
			"mapping_rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceBotMappingRuleSchema(),
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
