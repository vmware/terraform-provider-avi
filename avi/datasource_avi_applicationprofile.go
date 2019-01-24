/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviApplicationProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviApplicationProfileRead,
		Schema: map[string]*schema.Schema{
			"cloud_config_cksum": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"created_by": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_service_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDnsServiceApplicationProfileSchema(),
			},
			"dos_rl_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceDosRateLimitProfileSchema(),
			},
			"http_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHTTPApplicationProfileSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"preserve_client_ip": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"preserve_client_port": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"sip_service_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSipServiceApplicationProfileSchema(),
			},
			"tcp_app_profile": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceTCPApplicationProfileSchema(),
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
