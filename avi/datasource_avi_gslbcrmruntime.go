// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviGslbCRMRuntime() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbCRMRuntimeRead,
		Schema: map[string]*schema.Schema{
			"cluster_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceEventInfoSchema(),
			},
			"fds_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceFdsInfoSchema(),
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
			"remote_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceRemoteInfoSchema(),
			},
			"replication_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceReplicationPolicySchema(),
			},
			"site_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"status_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceOperationalStatusSchema(),
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
