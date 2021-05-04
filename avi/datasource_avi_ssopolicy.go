/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSSOPolicy() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSSOPolicyRead,
		Schema: map[string]*schema.Schema{
			"authentication_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAuthenticationPolicySchema(),
			},
			"authorization_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAuthorizationPolicySchema(),
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
			"type": {
				Type:     schema.TypeString,
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
