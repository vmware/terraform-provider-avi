// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

//nolint
func dataSourceAviIpamDnsProviderProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviIpamDnsProviderProfileRead,
		Schema: map[string]*schema.Schema{
			"allocate_ip_in_vrf": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpamDnsAwsProfileSchema(),
			},
			"azure_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpamDnsAzureProfileSchema(),
			},
			"custom_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpamDnsCustomProfileSchema(),
			},
			"gcp_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpamDnsGCPProfileSchema(),
			},
			"infoblox_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpamDnsInfobloxProfileSchema(),
			},
			"internal_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpamDnsInternalProfileSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oci_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpamDnsOCIProfileSchema(),
			},
			"openstack_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpamDnsOpenstackProfileSchema(),
			},
			"proxy_configuration": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProxyConfigurationSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tencent_profile": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceIpamDnsTencentProfileSchema(),
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
