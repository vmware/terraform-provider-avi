/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSecurityPolicyRead,
		Schema: map[string]*schema.Schema{
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_attacks": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDnsAttacksSchema(),
			},
			"dns_policy_index": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"network_security_policy_index": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"oper_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tcp_attacks": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTcpAttacksSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_attacks": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceUdpAttacksSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
