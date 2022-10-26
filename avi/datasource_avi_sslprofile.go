// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dhparam": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ec_named_curve": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_early_data": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enable_ssl_session_reuse": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_federated": {
				Type:     schema.TypeString,
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
			"prefer_client_cipher_ordering": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"send_close_notify": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"signature_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_rating": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSSLRatingSchema(),
			},
			"ssl_session_timeout": {
				Type:     schema.TypeString,
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
