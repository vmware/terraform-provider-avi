// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviCloudConnectorUser() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCloudConnectorUserRead,
		Schema: map[string]*schema.Schema{
			"azure_serviceprincipal": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAzureServicePrincipalCredentialsSchema(),
			},
			"azure_userpass": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAzureUserPassCredentialsSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"gcp_credentials": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceGCPCredentialsSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nsxt_credentials": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceNsxtCredentialsSchema(),
			},
			"oci_credentials": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceOCICredentialsSchema(),
			},
			"password": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"private_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"public_key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tencent_credentials": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTencentCredentialsSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vcenter_credentials": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceVCenterCredentialsSchema(),
			},
		},
	}
}
