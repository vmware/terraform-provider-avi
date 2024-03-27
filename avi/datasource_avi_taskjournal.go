// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviTaskJournal() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviTaskJournalRead,
		Schema: map[string]*schema.Schema{
			"errors": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceJournalErrorSchema(),
			},
			"image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceJournalInfoSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_cloud_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"operation": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"summary": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceJournalSummarySchema(),
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
