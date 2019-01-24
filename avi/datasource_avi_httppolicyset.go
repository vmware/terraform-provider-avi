/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviHTTPPolicySet() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHTTPPolicySetRead,
		Schema: map[string]*schema.Schema{
			"cloud_config_cksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http_request_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPRequestPolicySchema(),
			},
			"http_response_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPResponsePolicySchema(),
			},
			"http_security_policy": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPSecurityPolicySchema(),
			},
			"is_internal_policy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
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
