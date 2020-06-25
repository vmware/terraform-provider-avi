/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviLicenseLedgerDetails() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviLicenseLedgerDetailsRead,
		Schema: map[string]*schema.Schema{
			"escrow_infos": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceLicenseInfoSchema(),
			},
			"se_infos": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceLicenseInfoSchema(),
			},
			"tier_usages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceLicenseTierUsageSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
