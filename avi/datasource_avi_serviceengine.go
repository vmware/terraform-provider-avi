/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviServiceEngine() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviServiceEngineRead,
		Schema: map[string]*schema.Schema{
			"availability_zone": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"cloud_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"container_mode": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"container_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "CONTAINER_TYPE_HOST"},
			"controller_created": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"controller_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"data_vnics": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     ResourcevNICSchema(),
			},
			"enable_state": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "SE_STATE_ENABLED"},
			"flavor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"host_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hypervisor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"mgmt_vnic": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourcevNICSchema(),
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "VM name unknown"},
			"resources": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem:     ResourceSeResourcesSchema(),
			},
			"se_group_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"tenant_ref": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uuid": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}
