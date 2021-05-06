// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHealthMonitorRead,
		Schema: map[string]*schema.Schema{
			"allow_duplicate_monitors": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_quickstart": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"dns_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorDNSSchema(),
			},
			"external_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorExternalSchema(),
			},
			"failed_checks": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"http_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorHttpSchema(),
			},
			"https_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorHttpSchema(),
			},
			"is_federated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"monitor_port": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorRadiusSchema(),
			},
			"receive_timeout": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"send_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"sip_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorSIPSchema(),
			},
			"successful_checks": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tcp_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorTcpSchema(),
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"udp_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorUdpSchema(),
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
