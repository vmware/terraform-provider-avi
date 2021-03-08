/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviServiceEngine() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServiceEngineRead,
		Schema: map[string]*schema.Schema{
			"availability_zone": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"container_mode": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"container_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"controller_created": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"controller_ip": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"data_vnics": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourcevNICSchema(),
			},
			"enable_state": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"flavor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"host_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"hypervisor": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"mgmt_vnic": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourcevNICSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"resources": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSeResourcesSchema(),
			},
			"se_group_ref": {
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
