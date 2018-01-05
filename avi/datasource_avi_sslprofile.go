/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviSSLProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSSLProfileRead,
		Schema: map[string]*schema.Schema{
			"accepted_ciphers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "AES:3DES:RC4"},
			"accepted_versions": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceSSLVersionSchema(),
			},
			"cipher_enums": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dhparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"enable_ssl_session_reuse": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"prefer_client_cipher_ordering": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"send_close_notify": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"ssl_rating": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSSLRatingSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"ssl_session_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  86400,
			},
			"tags": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceTagSchema(),
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
