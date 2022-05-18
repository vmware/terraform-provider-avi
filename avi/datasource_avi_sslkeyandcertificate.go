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
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
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
			"enable_ocsp_stapling": {
				Type:     schema.TypeString,
				Computed: true,
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
			"import_key_to_hsm": {
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
			"ocsp_config": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceOCSPConfigSchema(),
			},
			"ocsp_error_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ocsp_responder_url_list_from_certs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"ocsp_response_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceOCSPResponseInfoSchema(),
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
