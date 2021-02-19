/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviVCenterServer() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviVCenterServerRead,
		Schema: map[string]*schema.Schema{
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"content_lib": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceContentLibConfigSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			"vcenter_credentials_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_url": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
