// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSystemLimitsSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"configpb_attributes": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceConfigPbAttributesSchema(),
		},
		"controller_limits": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceControllerLimitsSchema(),
		},
		"controller_sizes": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceControllerSizeSchema(),
		},
		"serviceengine_limits": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceServiceEngineLimitsSchema(),
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviSystemLimits() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSystemLimitsCreate,
		Read:   ResourceAviSystemLimitsRead,
		Update: resourceAviSystemLimitsUpdate,
		Delete: resourceAviSystemLimitsDelete,
		Schema: ResourceSystemLimitsSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSystemLimitsImporter,
		},
	}
}

func ResourceSystemLimitsImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSystemLimitsSchema()
	return ResourceImporter(d, m, "systemlimits", s)
}

func ResourceAviSystemLimitsRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemLimitsSchema()
	err := APIRead(d, meta, "systemlimits", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSystemLimitsCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemLimitsSchema()
	err := APICreateOrUpdate(d, meta, "systemlimits", s)
	if err == nil {
		err = ResourceAviSystemLimitsRead(d, meta)
	}
	return err
}

func resourceAviSystemLimitsUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemLimitsSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "systemlimits", s)
	if err == nil {
		err = ResourceAviSystemLimitsRead(d, meta)
	}
	return err
}

func resourceAviSystemLimitsDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "systemlimits")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
