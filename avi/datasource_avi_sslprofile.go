/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/helper/schema"

func dataSourceAviSSLProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSSLProfileRead,
		Schema: map[string]*schema.Schema{
			"accepted_ciphers": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"accepted_versions": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceSSLVersionSchema(),
			},
			"cipher_enums": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ciphersuites": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhparam": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_early_data": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_ssl_session_reuse": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"prefer_client_cipher_ordering": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"send_close_notify": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"ssl_rating": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSSLRatingSchema(),
			},
			"ssl_session_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tags": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceTagSchema(),
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
