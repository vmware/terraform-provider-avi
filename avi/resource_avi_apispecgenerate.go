// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceApiSpecGenerateSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"completed_events": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
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
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceApiSpecGenerateParamsSchema(),
		},
		"path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"progress": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
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
			Elem:     ResourceApiSpecGenerateStateSchema(),
		},
		"task_events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceTaskEventMapSchema(),
		},
		"tenant_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"total_events": {
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

func resourceAviApiSpecGenerate() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviApiSpecGenerateCreate,
		Read:   ResourceAviApiSpecGenerateRead,
		Update: resourceAviApiSpecGenerateUpdate,
		Delete: resourceAviApiSpecGenerateDelete,
		Schema: ResourceApiSpecGenerateSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceApiSpecGenerateImporter,
		},
	}
}

func ResourceApiSpecGenerateImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceApiSpecGenerateSchema()
	return ResourceImporter(d, m, "apispecgenerate", s)
}

func ResourceAviApiSpecGenerateRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiSpecGenerateSchema()
	err := APIRead(d, meta, "apispecgenerate", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviApiSpecGenerateCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiSpecGenerateSchema()
	err := APICreate(d, meta, "apispecgenerate", s)
	if err == nil {
		err = ResourceAviApiSpecGenerateRead(d, meta)
	}
	return err
}

func resourceAviApiSpecGenerateUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceApiSpecGenerateSchema()
	var err error
	err = APIUpdate(d, meta, "apispecgenerate", s)
	if err == nil {
		err = ResourceAviApiSpecGenerateRead(d, meta)
	}
	return err
}

func resourceAviApiSpecGenerateDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "apispecgenerate")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
