/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviHardwareSecurityModuleGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHardwareSecurityModuleGroupRead,
		Schema: map[string]*schema.Schema{
			"hsm": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,

				Elem: ResourceHardwareSecurityModuleSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
