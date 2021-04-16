/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviGslb() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviGslbRead,
		Schema: map[string]*schema.Schema{
			"async_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"clear_on_max_retries": {
				Type:     schema.TypeInt,
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
			"enable_config_by_members": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"error_resync_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"is_federated": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"leader_cluster_uuid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"maintenance_mode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replication_policy": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceReplicationPolicySchema(),
			},
			"send_interval": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"send_interval_prior_to_maintenance_mode": {
				Type:     schema.TypeInt,
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
				Type:     schema.TypeBool,
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
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}
