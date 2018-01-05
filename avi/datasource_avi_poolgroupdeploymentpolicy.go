/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPoolGroupDeploymentPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPoolGroupDeploymentPolicyRead,
		Schema: map[string]*schema.Schema{
			"auto_disable_old_prod_pools": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"evaluation_duration": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  300,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcePGDeploymentRuleSchema(),
			},
			"scheme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "BLUE_GREEN"},
			"target_test_traffic_ratio": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"test_traffic_ratio_rampup": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  100,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"webhook_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
