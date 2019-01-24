/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviAuthProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAuthProfileRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"http": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceAuthProfileHTTPClientParamsSchema(),
			},
			"ldap": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceLdapAuthSettingsSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"saml": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSamlSettingsSchema(),
			},
			"tacacs_plus": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTacacsPlusAuthSettingsSchema(),
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
