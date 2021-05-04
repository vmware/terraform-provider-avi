/*
* Copyright (c) 2017. Avi Networks.
* Author: Gaurav Rastogi (grastogi@avinetworks.com)
*
 */
package avi

import "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

func dataSourceAviUpgradeStatusInfo() *schema.Resource {
	return &schema.Resource{
		Read: ResourceAviUpgradeStatusInfoRead,
		Schema: map[string]*schema.Schema{
			"after_reboot_rollback_fnc": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"after_reboot_task_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"clean": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"duration": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"enable_patch_rollback": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"enable_rollback": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"end_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"enqueue_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"node_type": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"obj_cloud_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceUpgradeOpsParamSchema(),
			},
			"patch_image_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"patch_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourcePatchDataSchema(),
			},
			"patch_reboot": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"patch_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prev_image_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"prev_patch_image_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_patch_image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_patch_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourcePatchDataSchema(),
			},
			"previous_patch_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"previous_version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"progress": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"se_patch_image_path": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_patch_image_ref": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"se_upgrade_events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceSeUpgradeEventsSchema(),
			},
			"seg_params": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceUpgradeOpsParamSchema(),
			},
			"seg_status": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceSeGroupStatusSchema(),
			},
			"start_time": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"state": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem:     ResourceUpgradeOpsStateSchema(),
			},
			"system": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"tasks_completed": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"tenant_ref": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"total_tasks": {
				Type:     schema.TypeInt,
				Computed: true,
			},
			"upgrade_events": {
				Type:     schema.TypeList,
				Computed: true,
				Elem:     ResourceEventMapSchema(),
			},
			"upgrade_ops": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"uuid": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}
