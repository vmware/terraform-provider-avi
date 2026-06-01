// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviCertJwtStore() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviCertJwtStoreRead,
		Schema: map[string]*schema.Schema{
			"configpb_attributes": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceConfigPbAttributesSchema(),
			},
			"jwt": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_passphrase": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"kid": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"last_rotated_at": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceTimeStampSchema(),
			},
			"public_key_algorithm": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": {
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
