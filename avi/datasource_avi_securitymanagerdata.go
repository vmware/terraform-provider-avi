/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSecurityManagerData() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSecurityManagerDataRead,
		Schema: map[string]*schema.Schema{
			"app_learning_info": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDbAppLearningInfoSchema(),
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
