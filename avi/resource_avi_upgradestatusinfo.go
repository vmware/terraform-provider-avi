// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Mozilla Public License 2.0

package avi

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"log"
)

func ResourceUpgradeStatusInfoSchema() map[string]*schema.Schema {
	return map[string]*schema.Schema{
		"after_reboot_rollback_fnc": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"after_reboot_task_name": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"clean": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"duration": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateInteger,
		},
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
		"enqueue_time": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"fips_mode": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"history": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceOpsHistorySchema(),
		},
		"image_path": {
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
		"params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceUpgradeOpsParamSchema(),
		},
		"patch_image_path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"patch_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePatchDataSchema(),
		},
		"patch_reboot": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
		},
		"patch_version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"prev_image_path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"prev_patch_image_path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"previous_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"previous_patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"previous_patch_list": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourcePatchDataSchema(),
		},
		"previous_patch_version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"previous_version": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"progress": {
			Type:         schema.TypeString,
			Optional:     true,
			Default:      "0",
			ValidateFunc: validateInteger,
		},
		"reason": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_patch_image_path": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_patch_image_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"se_upgrade_events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceSeUpgradeEventsSchema(),
		},
		"seg_params": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceUpgradeOpsParamSchema(),
		},
		"seg_status": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceSeGroupStatusSchema(),
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
		"statediff_ref": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"system": {
			Type:         schema.TypeString,
			Optional:     true,
			Computed:     true,
			ValidateFunc: validateBool,
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
		"upgrade_events": {
			Type:     schema.TypeList,
			Optional: true,
			Elem:     ResourceEventMapSchema(),
		},
		"upgrade_ops": {
			Type:     schema.TypeString,
			Optional: true,
			Computed: true,
		},
		"upgrade_readiness": {
			Type:     schema.TypeSet,
			Optional: true,
			Computed: true,
			Elem:     ResourceUpgradeReadinessCheckObjSchema(),
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

func resourceAviUpgradeStatusInfo() *schema.Resource {
	return &schema.Resource{
		Create: resourceAviUpgradeStatusInfoCreate,
		Read:   ResourceAviUpgradeStatusInfoRead,
		Update: resourceAviUpgradeStatusInfoUpdate,
		Delete: resourceAviUpgradeStatusInfoDelete,
		Schema: ResourceUpgradeStatusInfoSchema(),
		Importer: &schema.ResourceImporter{
			State: ResourceUpgradeStatusInfoImporter,
		},
	}
}

func ResourceUpgradeStatusInfoImporter(d *schema.ResourceData, m interface{}) ([]*schema.ResourceData, error) {
	s := ResourceUpgradeStatusInfoSchema()
	return ResourceImporter(d, m, "upgradestatusinfo", s)
}

func ResourceAviUpgradeStatusInfoRead(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusInfoSchema()
	err := APIRead(d, meta, "upgradestatusinfo", s)
	if err != nil {
		log.Printf("[ERROR] in reading object %v\n", err)
	}
	return err
}

func resourceAviUpgradeStatusInfoCreate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusInfoSchema()
	err := APICreateOrUpdate(d, meta, "upgradestatusinfo", s)
	if err == nil {
		err = ResourceAviUpgradeStatusInfoRead(d, meta)
	}
	return err
}

func resourceAviUpgradeStatusInfoUpdate(d *schema.ResourceData, meta interface{}) error {
	s := ResourceUpgradeStatusInfoSchema()
	var err error
	err = APICreateOrUpdate(d, meta, "upgradestatusinfo", s)
	if err == nil {
		err = ResourceAviUpgradeStatusInfoRead(d, meta)
	}
	return err
}

func resourceAviUpgradeStatusInfoDelete(d *schema.ResourceData, meta interface{}) error {
	var err error
	if APIDeleteSystemDefaultCheck(d) {
		return nil
	}
	err = APIDelete(d, meta, "upgradestatusinfo")
	if err != nil {
		log.Printf("[ERROR] in deleting object %v\n", err)
	}
	return err
}
