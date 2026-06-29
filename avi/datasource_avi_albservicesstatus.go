// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviALBServicesStatus() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviALBServicesStatusRead,
		Schema: map[string]*schema.Schema{
			"asset_details": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceALBServicesAssetDetailsSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"connected_at": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"connectivity_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"error": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"registration_status": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"services_health": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceServiceHealthSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_status": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcePulseServicesTenantStatusSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
