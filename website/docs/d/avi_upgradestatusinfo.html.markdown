############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_upgradestatusinfo"
sidebar_current: "docs-avi-datasource-upgradestatusinfo"
description: |-
  Get information of Avi UpgradeStatusInfo.
---

# avi_upgradestatusinfo

This data source is used to to get avi_upgradestatusinfo objects.

## Example Usage

```hcl
data "avi_upgradestatusinfo" "foo_upgradestatusinfo" {
    uuid = "upgradestatusinfo-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search UpgradeStatusInfo by name.
* `uuid` - (Optional) Search UpgradeStatusInfo by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `after_reboot_rollback_fnc` - Backward compatible abort function name.
* `after_reboot_task_name` - Backward compatible task dict name.
* `clean` - Flag for clean installation.
* `duration` - Duration of upgrade operation in seconds.
* `enable_patch_rollback` - Check if the patch rollback is possible on this node.
* `enable_rollback` - Check if the rollback is possible on this node.
* `end_time` - End time of upgrade operation.
* `enqueue_time` - Enqueue time of upgrade operation.
* `image_path` - Image path of current base image.
* `image_ref` - Image uuid for identifying the current base image.
* `name` - Name of the system such as cluster name, se group name and se name.
* `node_type` - Type of the system such as controller_cluster, se_group or se.
* `obj_cloud_ref` - Cloud that this object belongs to.
* `params` - Parameters associated with the upgrade operation.
* `patch_image_path` - Image path of current patch image.
* `patch_image_ref` - Image uuid for identifying the current patch.example  base-image is 18.2.6 and a patch 6p1 is applied, then this field will indicate the 6p1 value.
* `patch_list` - List of patches applied to this node.
* `patch_reboot` - Flag for patch op with reboot.
* `patch_version` - Current patch version applied to this node.
* `prev_image_path` - Image path of previous base image.
* `prev_patch_image_path` - Image path of previous patch image.
* `previous_image_ref` - Image uuid for identifying previous base image.example  base-image was 18.2.5 and an upgrade was done to 18.2.6, then this field will indicate the 18.2.5 value.
* `previous_patch_image_ref` - Image uuid for identifying previous patch.example  base-image was 18.2.6 with a patch 6p1.
* `previous_patch_list` - List of patches applied to this node on previous major version.
* `previous_patch_version` - Previous patch version applied to this node.example  base-image was 18.2.6 with a patch 6p1.
* `previous_version` - Previous version prior to upgrade.example  base-image was 18.2.5 and an upgrade was done to 18.2.6, then this field will indicate the 18.2.5 value.
* `progress` - Upgrade operations progress which holds value between 0-100.
* `se_patch_image_path` - Image path of se patch image.(required in case of reimage and upgrade + patch).
* `se_patch_image_ref` - Image uuid for identifying the current se patch required in case of system upgrade(re-image) with se patch.
* `se_upgrade_events` - Serviceenginegroup upgrade errors.
* `seg_params` - Se_patch may be different from the controller_patch.
* `seg_status` - Detailed segroup status.
* `start_time` - Start time of upgrade operation.
* `state` - Current status of the upgrade operation.
* `system` - Flag is set only in the cluster if the upgrade is initiated as a system-upgrade.
* `tasks_completed` - Completed set of tasks in the upgrade operation.
* `tenant_ref` - Tenant that this object belongs to.
* `total_tasks` - Total number of tasks in the upgrade operation.
* `upgrade_events` - Events performed for upgrade operation.
* `upgrade_ops` - Upgrade operations requested.
* `uuid` - Uuid identifier for the system such as cluster, se group and se.
* `version` - Current base image applied to this node.

