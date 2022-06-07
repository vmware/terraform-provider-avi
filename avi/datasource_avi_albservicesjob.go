// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviALBServicesJob() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviALBServicesJobRead,
		Schema: map[string]*schema.Schema{
			"command": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"end_time": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"params": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceALBServicesJobParamSchema(),
			},
			"pulse_job_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pulse_sync_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"result": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"start_time": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"token": {
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
