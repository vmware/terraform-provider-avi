// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceLocalWorkerFdsVersionSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Default:  "default",
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"timeline": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"version": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
	}
}

func resourceAviLocalWorkerFdsVersion() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviLocalWorkerFdsVersionCreate,
		Read:   ResourceAviLocalWorkerFdsVersionRead,
		Update: resourceAviLocalWorkerFdsVersionUpdate,
		Delete: resourceAviLocalWorkerFdsVersionDelete,
		Schema: ResourceLocalWorkerFdsVersionSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceLocalWorkerFdsVersionImporter,
		},
	}
}

func ResourceLocalWorkerFdsVersionImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceLocalWorkerFdsVersionSchema()
	return ResourceImporter(d, m, "localworkerfdsversion", s)
}

func ResourceAviLocalWorkerFdsVersionRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLocalWorkerFdsVersionSchema()
	err := APIRead(d, meta, "localworkerfdsversion", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviLocalWorkerFdsVersionCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLocalWorkerFdsVersionSchema()
	err := APICreate(d, meta, "localworkerfdsversion", s)
	if err == nil {
		err = ResourceAviLocalWorkerFdsVersionRead(d, meta)
	}
	return err
}

func resourceAviLocalWorkerFdsVersionUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceLocalWorkerFdsVersionSchema()
	var err error
	err = APIUpdate(d, meta, "localworkerfdsversion", s)
	if err == nil {
		err = ResourceAviLocalWorkerFdsVersionRead(d, meta)
	}
	return err
}

func resourceAviLocalWorkerFdsVersionDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "localworkerfdsversion")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
