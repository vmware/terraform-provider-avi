// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviGslb() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbRead,
		Schema: map[string]*schema.Schema{
			"async_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"clear_on_max_retries": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_ip_addr_group": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceGslbClientIpAddrGroupSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"dns_configs": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceDNSConfigSchema(),
			},
			"error_resync_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_federated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"leader_cluster_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"send_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"send_interval_prior_to_maintenance_mode": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sites": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceGslbSiteSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_scoped": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"third_party_sites": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceGslbThirdPartySiteSchema(),
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
