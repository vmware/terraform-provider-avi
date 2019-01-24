/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviHealthMonitor() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviHealthMonitorRead,
		Schema: map[string]*schema.Schema{
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"dns_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorDNSSchema(),
			},
			"external_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorExternalSchema(),
			},
			"failed_checks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"http_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorHttpSchema(),
			},
			"https_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorHttpSchema(),
			},
			"is_federated": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"monitor_port": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"receive_timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  4,
			},
			"send_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  10,
			},
			"sip_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorSIPSchema(),
			},
			"successful_checks": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  2,
			},
			"tcp_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorTcpSchema(),
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"udp_monitor": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceHealthMonitorUdpSchema(),
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
