/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviPKIProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviPKIProfileRead,
		Schema: map[string]*schema.Schema{
			"ca_certs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSSLCertificateSchema(),
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"crl_check": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"crls": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCRLSchema(),
			},
			"ignore_peer_chain": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"is_federated": &schema.Schema{
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
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"validate_only_leaf_crl": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}
