// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviAkoAmkoCluster() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviAkoAmkoClusterRead,
		Schema: map[string]*schema.Schema{
			"cloud_config_cksum": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"cloud_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cluster_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"created_by": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"deployment_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAkoAmkoClusterDeploymentInfoSchema(),
			},
			"metadata": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAkoAmkoClusterMetadataSchema(),
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
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
			"version_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceAkoAmkoClusterVersionInfoSchema(),
			},
		},
	}
}
