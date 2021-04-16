/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviLabelGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviLabelGroupRead,
		Schema: map[string]*schema.Schema{
			"labels": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleMatchOperationMatchLabelSchema(),
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
