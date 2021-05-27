// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

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
