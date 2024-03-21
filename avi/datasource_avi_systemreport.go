// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviSystemReport() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviSystemReportRead,
		Schema: map[string]*schema.Schema{
			"archive_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_patch_image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"downloadable": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceReportEventSchema(),
			},
			"image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"readiness_reports": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceReportDetailSchema(),
			},
			"se_patch_image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceReportOpsStateSchema(),
			},
			"summary": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceReportSummarySchema(),
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
