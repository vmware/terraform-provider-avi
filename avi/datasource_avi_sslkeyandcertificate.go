/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSSLKeyAndCertificate() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSSLKeyAndCertificateRead,
		Schema: map[string]*schema.Schema{
			"ca_certs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCertificateAuthoritySchema(),
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLCertificateSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"certificate_base64": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"certificate_management_profile_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dynamic_params": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCustomParamsSchema(),
			},
			"enckey_base64": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enckey_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"format": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_PEM"},
			"hardwaresecuritymodulegroup_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"key_base64": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"key_params": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLKeyParamsSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"key_passphrase": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SSL_CERTIFICATE_FINISHED"},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
