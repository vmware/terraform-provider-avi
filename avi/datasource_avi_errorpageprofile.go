/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviErrorPageProfile() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviErrorPageProfileRead,
		Schema: map[string]*schema.Schema{
			"app_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VS Name"},
			"company_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Avi Networks"},
			"error_pages": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourceErrorPageSchema(),
			},
			"host_name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Host Header"},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
