/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviTrafficCloneProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviTrafficCloneProfileRead,
		Schema: map[string]*schema.Schema{
			"clone_servers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceCloneServerSchema(),
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "/api/cloud?name=Default-Cloud",
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
			"tenant_ref": &schema.Schema{
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
