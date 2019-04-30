/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviCloudConnectorUser() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCloudConnectorUserRead,
		Schema: map[string]*schema.Schema{
			"azure_serviceprincipal": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureServicePrincipalCredentialsSchema(),
			},
			"azure_userpass": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAzureUserPassCredentialsSchema(),
			},
			"gcp_credentials": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceGCPCredentialsSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_credentials": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceOCICredentialsSchema(),
			},
			"password": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"private_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"public_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tencent_credentials": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTencentCredentialsSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
