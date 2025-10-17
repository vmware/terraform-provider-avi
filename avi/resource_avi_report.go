// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
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
		"filename": {
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
		"pre_check": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceReadinessCheckObjSchema(),
		},
		"progress": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"request": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceReportGenerationRequestSchema(),
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
			Elem:     ResourceReportGenStateSchema(),
		},
		"tasks": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceTaskEventMapSchema(),
		},
		"tasks_completed": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
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
	}
}

func resourceAviReport() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviReportCreate,
		Read:   ResourceAviReportRead,
		Update: resourceAviReportUpdate,
		Delete: resourceAviReportDelete,
		Schema: ResourceReportSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceReportImporter,
		},
	}
}

func ResourceReportImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceReportSchema()
	return ResourceImporter(d, m, "report", s)
}

func ResourceAviReportRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceReportSchema()
	err := APIRead(d, meta, "report", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviReportCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceReportSchema()
	err := APICreate(d, meta, "report", s)
	if err == nil {
		err = ResourceAviReportRead(d, meta)
	}
	return err
}

func resourceAviReportUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceReportSchema()
	var err error
	err = APIUpdate(d, meta, "report", s)
	if err == nil {
		err = ResourceAviReportRead(d, meta)
	}
	return err
}

func resourceAviReportDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "report")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
