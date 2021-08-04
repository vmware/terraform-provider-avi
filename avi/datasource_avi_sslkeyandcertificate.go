// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSSLKeyAndCertificate() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSSLKeyAndCertificateRead,
		Schema: map[string]*schema.Schema{
			"ca_certs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCertificateAuthoritySchema(),
			},
			"certificate": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSSLCertificateSchema(),
			},
			"certificate_base64": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate_management_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dynamic_params": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCustomParamsSchema(),
			},
			"enckey_base64": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enckey_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"format": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hardwaresecuritymodulegroup_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_base64": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSSLKeyParamsSchema(),
			},
			"key_passphrase": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": {
				Type:     schema.TypeString,
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
