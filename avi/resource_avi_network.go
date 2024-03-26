// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceNetworkSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"attrs": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceKeyValueSchema(),
		},
		"cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"configured_subnets": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSubnetSchema(),
		},
		"dhcp_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"exclude_discovered_subnets": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"ip6_autocfg_enabled": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"markers": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceRoleFilterMatchLabelSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Required: true,
		},
		"synced_from_se": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
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
		"vcenter_dvs": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "true",
			ValidateFunc: validateBool,
		},
		"vimgrnw_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"vrf_context_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviNetwork() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviNetworkCreate,
		Read:   ResourceAviNetworkRead,
		Update: resourceAviNetworkUpdate,
		Delete: resourceAviNetworkDelete,
		Schema: ResourceNetworkSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceNetworkImporter,
		},
	}
}

func ResourceNetworkImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceNetworkSchema()
	return ResourceImporter(d, m, "network", s)
}

func ResourceAviNetworkRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSchema()
	err := APIRead(d, meta, "network", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviNetworkCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSchema()
	err := APICreate(d, meta, "network", s)
	if err == nil {
		err = ResourceAviNetworkRead(d, meta)
	}
	return err
}

func resourceAviNetworkUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceNetworkSchema()
	var err error
	err = APIUpdate(d, meta, "network", s)
	if err == nil {
		err = ResourceAviNetworkRead(d, meta)
	}
	return err
}

func resourceAviNetworkDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "network")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
