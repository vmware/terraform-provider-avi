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
				Set: func(v interface{}) int {
					return 0
				},
			},
			"azure_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsAzureProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"custom_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsCustomProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"gcp_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsGCPProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"infoblox_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsInfobloxProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"internal_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsInternalProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"openstack_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceIpamDnsOpenstackProfileSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"proxy_configuration": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceProxyConfigurationSchema(),
				Set: func(v interface{}) int {
					return 0
				},
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": &schema.Schema{
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
