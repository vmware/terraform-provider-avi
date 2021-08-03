// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHealthMonitorRead,
		Schema: map[string]*schema.Schema{
			"allow_duplicate_monitors": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"authentication": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorAuthInfoSchema(),
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"description": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"disable_quickstart": {
				Type:     schema.TypeString,
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
				Type:     schema.TypeString,
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
			"imap_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorImapSchema(),
			},
			"imaps_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorImapSchema(),
			},
			"is_federated": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"markers": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceRoleFilterMatchLabelSchema(),
			},
			"monitor_port": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pop3_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorPop3Schema(),
			},
			"pop3s_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorPop3Schema(),
			},
			"radius_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorRadiusSchema(),
			},
			"receive_timeout": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"send_interval": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"sip_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorSIPSchema(),
			},
			"smtp_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorSmtpSchema(),
			},
			"smtps_monitor": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceHealthMonitorSmtpSchema(),
			},
			"successful_checks": {
				Type:     schema.TypeString,
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
