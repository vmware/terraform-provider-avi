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
			"clone_servers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceCloneServerSchema(),
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
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
