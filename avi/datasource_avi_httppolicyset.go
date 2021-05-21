/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviHTTPPolicySet() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHTTPPolicySetRead,
		Schema: map[string]*schema.Schema{
			"cloud_config_cksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"geo_db_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_request_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTPRequestPolicySchema(),
			},
			"http_response_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTPResponsePolicySchema(),
			},
			"http_security_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTPSecurityPolicySchema(),
			},
			"ip_reputation_db_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_internal_policy": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
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
