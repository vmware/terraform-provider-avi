/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviRole() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviRoleRead,
		Schema: map[string]*schema.Schema{
			"allow_unlabelled_access": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"filters": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"privileges": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourcePermissionSchema(),
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
