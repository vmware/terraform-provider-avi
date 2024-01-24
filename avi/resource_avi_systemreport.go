// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceSystemReportSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"archive_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"controller_patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"downloadable": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceReportEventSchema(),
		},
		"image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"readiness_reports": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceReportDetailSchema(),
		},
		"se_patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"state": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceReportOpsStateSchema(),
		},
		"summary": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceReportSummarySchema(),
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
	}
}

func resourceAviSystemReport() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviSystemReportCreate,
		Read:   ResourceAviSystemReportRead,
		Update: resourceAviSystemReportUpdate,
		Delete: resourceAviSystemReportDelete,
		Schema: ResourceSystemReportSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceSystemReportImporter,
		},
	}
}

func ResourceSystemReportImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceSystemReportSchema()
	return ResourceImporter(d, m, "systemreport", s)
}

func ResourceAviSystemReportRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemReportSchema()
	err := APIRead(d, meta, "systemreport", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviSystemReportCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemReportSchema()
	err := APICreateOrUpdate(d, meta, "systemreport", s)
	if err == nil {
		err = ResourceAviSystemReportRead(d, meta)
	}
	return err
}

func resourceAviSystemReportUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceSystemReportSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "systemreport", s)
	if err == nil {
		err = ResourceAviSystemReportRead(d, meta)
	}
	return err
}

func resourceAviSystemReportDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "systemreport")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
