// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviRmCloudOpsProto() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviRmCloudOpsProtoRead,
		Schema: map[string]*schema.Schema{
			"last_queried_se_creation_limit": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"pending_se_creation_count": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"pending_vnic_op_count": {
				Type:     schema.TypeString,
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
