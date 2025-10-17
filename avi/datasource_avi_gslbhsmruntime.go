// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviGslbHSMRuntime() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbHSMRuntimeRead,
		Schema: map[string]*schema.Schema{
			"cluster_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enabled": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceEventInfoSchema(),
			},
			"local_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceLocalInfoSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"obj_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"oper_status": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
			},
			"remote_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRemoteInfoSchema(),
			},
			"send_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_name": {
				Type:     schema.TypeString,
				Computed: true,
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
