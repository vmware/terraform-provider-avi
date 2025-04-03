// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviGslbSMRuntime() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbSMRuntimeRead,
		Schema: map[string]*schema.Schema{
			"cluster_leader": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cluster_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_configs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDNSConfigSchema(),
			},
			"dns_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceGslbDnsInfoSchema(),
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
			"health_monitor_info": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"leader_cluster_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"member_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"num_of_retries": {
				Type:     schema.TypeString,
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
			"role": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"site_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sw_version": {
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
			"view_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
