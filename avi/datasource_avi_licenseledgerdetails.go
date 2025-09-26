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
			"se_group_infos": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceSeGroupInfoSchema(),
			},
			"se_infos": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceLicenseInfoSchema(),
			},
			"tenant_infos": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceLicenseReservationInfoSchema(),
			},
			"tier_usages": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceLicenseTierUsageSchema(),
			},
			"total_licenses_reserved": {
				Type:     schema.TypeString,
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
