// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviMemoryBalancerRequest() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviMemoryBalancerRequestRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"controller_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceControllerInfoSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"process_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceProcessInfoSchema(),
			},
			"process_instance": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"timestamp": {
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
