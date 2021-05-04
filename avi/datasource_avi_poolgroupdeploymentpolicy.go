/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviPoolGroupDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPoolGroupDeploymentPolicyRead,
		Schema: map[string]*schema.Schema{
			"auto_disable_old_prod_pools": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"evaluation_duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rules": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourcePGDeploymentRuleSchema(),
			},
			"scheme": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"target_test_traffic_ratio": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"test_traffic_ratio_rampup": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"webhook_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
