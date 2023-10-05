// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceUpgradeStatusSummarySchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"enable_patch_rollback": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"enable_rollback": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "false",
			ValidateFunc: validateBool,
		},
		"end_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
		"node_type": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"obj_cloud_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
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
			Elem:     ResourceUpgradeOpsStateSchema(),
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
		"upgrade_ops": {
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
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
	}
}

func resourceAviUpgradeStatusSummary() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUpgradeStatusSummaryCreate,
		Read:   ResourceAviUpgradeStatusSummaryRead,
		Update: resourceAviUpgradeStatusSummaryUpdate,
		Delete: resourceAviUpgradeStatusSummaryDelete,
		Schema: ResourceUpgradeStatusSummarySchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceUpgradeStatusSummaryImporter,
		},
	}
}

func ResourceUpgradeStatusSummaryImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUpgradeStatusSummarySchema()
	return ResourceImporter(d, m, "upgradestatussummary", s)
}

func ResourceAviUpgradeStatusSummaryRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusSummarySchema()
	err := APIRead(d, meta, "upgradestatussummary", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUpgradeStatusSummaryCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusSummarySchema()
	err := APICreateOrUpdate(d, meta, "upgradestatussummary", s)
	if err == nil {
		err = ResourceAviUpgradeStatusSummaryRead(d, meta)
	}
	return err
}

func resourceAviUpgradeStatusSummaryUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusSummarySchema()
	var err error
	err = APICreateOrUpdate(d, meta, "upgradestatussummary", s)
	if err == nil {
		err = ResourceAviUpgradeStatusSummaryRead(d, meta)
	}
	return err
}

func resourceAviUpgradeStatusSummaryDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "upgradestatussummary")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
