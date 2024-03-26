// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceNsxtSegmentRuntimeSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"dhcp6_ranges": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"dhcp_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"dhcp_ranges": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"nw_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"nw_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"opaque_network_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"origin_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"security_only_nsxt": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"segment_gw": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"segment_gw6": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"segment_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"segname": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"subnet": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"subnet6": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"tier1_id": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vlan_ids": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"vrf_context_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviNsxtSegmentRuntime() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNsxtSegmentRuntimeCreate,
		Read:   ResourceAviNsxtSegmentRuntimeRead,
		Update: resourceAviNsxtSegmentRuntimeUpdate,
		Delete: resourceAviNsxtSegmentRuntimeDelete,
		Schema: ResourceNsxtSegmentRuntimeSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNsxtSegmentRuntimeImporter,
		},
	}
}

func ResourceNsxtSegmentRuntimeImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNsxtSegmentRuntimeSchema()
	return ResourceImporter(d, m, "nsxtsegmentruntime", s)
}

func ResourceAviNsxtSegmentRuntimeRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNsxtSegmentRuntimeSchema()
	err := APIRead(d, meta, "nsxtsegmentruntime", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNsxtSegmentRuntimeCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNsxtSegmentRuntimeSchema()
	err := APICreateOrUpdate(d, meta, "nsxtsegmentruntime", s)
	if err == nil {
		err = ResourceAviNsxtSegmentRuntimeRead(d, meta)
	}
	return err
}

func resourceAviNsxtSegmentRuntimeUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNsxtSegmentRuntimeSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "nsxtsegmentruntime", s)
	if err == nil {
		err = ResourceAviNsxtSegmentRuntimeRead(d, meta)
	}
	return err
}

func resourceAviNsxtSegmentRuntimeDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "nsxtsegmentruntime")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
