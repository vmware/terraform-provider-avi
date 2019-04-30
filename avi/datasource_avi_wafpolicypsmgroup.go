/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviWafPolicyPSMGroup() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviWafPolicyPSMGroupRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"hit_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_ACTION_NO_OP"},
			"is_learning_group": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"locations": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceWafPSMLocationSchema(),
			},
			"miss_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "WAF_ACTION_NO_OP"},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
