// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceInventoryFaultConfigSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"controller_faults": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceControllerFaultsSchema(),
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"serviceengine_faults": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceServiceengineFaultsSchema(),
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
		"virtualservice_faults": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceVirtualserviceFaultsSchema(),
		},
	}
}

func resourceAviInventoryFaultConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviInventoryFaultConfigCreate,
		Read:   ResourceAviInventoryFaultConfigRead,
		Update: resourceAviInventoryFaultConfigUpdate,
		Delete: resourceAviInventoryFaultConfigDelete,
		Schema: ResourceInventoryFaultConfigSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceInventoryFaultConfigImporter,
		},
	}
}

func ResourceInventoryFaultConfigImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceInventoryFaultConfigSchema()
	return ResourceImporter(d, m, "inventoryfaultconfig", s)
}

func ResourceAviInventoryFaultConfigRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceInventoryFaultConfigSchema()
	err := APIRead(d, meta, "inventoryfaultconfig", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviInventoryFaultConfigCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceInventoryFaultConfigSchema()
	err := APICreate(d, meta, "inventoryfaultconfig", s)
	if err == nil {
		err = ResourceAviInventoryFaultConfigRead(d, meta)
	}
	return err
}

func resourceAviInventoryFaultConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceInventoryFaultConfigSchema()
	var err error
	err = APIUpdate(d, meta, "inventoryfaultconfig", s)
	if err == nil {
		err = ResourceAviInventoryFaultConfigRead(d, meta)
	}
	return err
}

func resourceAviInventoryFaultConfigDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "inventoryfaultconfig")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
