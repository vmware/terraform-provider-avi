/*
 * Copyright (c) 2017. Avi Networks.
 * Author: Gaurav Rastogi (grastogi@avinetworks.com)
 *
 */
package avi

import "github.com/hashicorp/terraform/helper/schema"

func dataSourceAviRmCloudOpsProto() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviRmCloudOpsProtoRead,
		Schema: map[string]*schema.Schema{
			"last_queried_se_creation_limit": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pending_se_creation_count": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"pending_vnic_op_count": {
				Type:     schema.TypeInt,
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
