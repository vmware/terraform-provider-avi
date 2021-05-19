/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviBotIPReputationTypeMapping() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviBotIPReputationTypeMappingRead,
		Schema: map[string]*schema.Schema{
			"ip_reputation_mappings": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIPReputationTypeMappingSchema(),
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
