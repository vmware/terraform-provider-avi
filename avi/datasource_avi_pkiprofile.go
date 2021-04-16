/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviPKIProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPKIProfileRead,
		Schema: map[string]*schema.Schema{
			"ca_certs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceSSLCertificateSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"crl_check": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"crls": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCRLSchema(),
			},
			"ignore_peer_chain": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_federated": {
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
			"validate_only_leaf_crl": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}
