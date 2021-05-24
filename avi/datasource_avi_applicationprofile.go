// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviApplicationProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApplicationProfileRead,
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
			"dns_service_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDnsServiceApplicationProfileSchema(),
			},
			"dos_rl_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceDosRateLimitProfileSchema(),
			},
			"http_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHTTPApplicationProfileSchema(),
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
			"preserve_client_ip": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"preserve_client_port": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"preserve_dest_ip_port": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"sip_service_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSipServiceApplicationProfileSchema(),
			},
			"tcp_app_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTCPApplicationProfileSchema(),
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
