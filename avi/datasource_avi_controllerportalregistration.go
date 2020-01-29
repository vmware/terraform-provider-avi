/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviControllerPortalRegistration() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviControllerPortalRegistrationRead,
		Schema: map[string]*schema.Schema{
			"asset": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerPortalAssetSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"portal_auth": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerPortalAuthSchema(),
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
