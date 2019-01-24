/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSecurityPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSecurityPolicyRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_attacks": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsAttacksSchema(),
			},
			"dns_policy_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"network_security_policy_index": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  0,
			},
			"oper_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "DETECTION"},
			"tcp_attacks": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTcpAttacksSchema(),
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"udp_attacks": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceUdpAttacksSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
