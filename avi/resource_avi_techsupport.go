// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceTechSupportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"case_number": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"description": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"duration": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"end_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"errors": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
		"level": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"node": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"obj_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"obj_uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"output": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTechSupportParamsSchema(),
		},
		"progress": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"size": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateFloat,
		},
		"start_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceTechSupportStateSchema(),
		},
		"tasks": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceTechSupportEventMapSchema(),
		},
		"tasks_completed": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"techsupport_readiness": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceReadinessCheckObjSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"total_tasks": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
		"uuid": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"warnings": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},
	}
}

func resourceAviTechSupport() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviTechSupportCreate,
		Read:   ResourceAviTechSupportRead,
		Update: resourceAviTechSupportUpdate,
		Delete: resourceAviTechSupportDelete,
		Schema: ResourceTechSupportSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceTechSupportImporter,
		},
	}
}

func ResourceTechSupportImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceTechSupportSchema()
	return ResourceImporter(d, m, "techsupport", s)
}

func ResourceAviTechSupportRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTechSupportSchema()
	err := APIRead(d, meta, "techsupport", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviTechSupportCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTechSupportSchema()
	err := APICreate(d, meta, "techsupport", s)
	if err == nil {
		err = ResourceAviTechSupportRead(d, meta)
	}
	return err
}

func resourceAviTechSupportUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceTechSupportSchema()
	var err error
	err = APIUpdate(d, meta, "techsupport", s)
	if err == nil {
		err = ResourceAviTechSupportRead(d, meta)
	}
	return err
}

func resourceAviTechSupportDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "techsupport")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
