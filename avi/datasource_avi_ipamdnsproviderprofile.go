/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviIpamDnsProviderProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviIpamDnsProviderProfileRead,
		Schema: map[string]*schema.Schema{
			"allocate_ip_in_vrf": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"aws_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsAwsProfileSchema(),
			},
			"azure_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsAzureProfileSchema(),
			},
			"custom_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsCustomProfileSchema(),
			},
			"gcp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsGCPProfileSchema(),
			},
			"infoblox_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsInfobloxProfileSchema(),
			},
			"internal_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsInternalProfileSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsOCIProfileSchema(),
			},
			"openstack_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsOpenstackProfileSchema(),
			},
			"proxy_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProxyConfigurationSchema(),
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tencent_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsTencentProfileSchema(),
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
