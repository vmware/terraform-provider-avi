// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSessionKeyForwarder() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSessionKeyForwarderRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"enable": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ip_ports": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceIpAddrPortSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pki_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_key_and_certificate_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"ssl_profile_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_mgmt": {
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
