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

* `duration` - Duration of upgrade operations in seconds.
* `enable_patch_rollback` - Check if the patch rollback is possible on this node.
* `enable_rollback` - Check if the rollback is possible on this node.
* `end_time` - End time of upgrade operations.
* `enqueue_time` - Enqueue time of upgrade operations.
* `image_ref` - Image uuid for identifying the current base image.
* `name` - Name of the system such as cluster name, se group name and se name.
* `node_type` - Type of the system such as controller_cluster, se_group or se.
* `obj_cloud_ref` - Cloud that this object belongs to.
* `params` - Parameters associated with the current upgrade ops.
* `patch_image_ref` - Image uuid for identifying the current patch.
* `patch_version` - Current patch version applied to this node.
* `previous_image_ref` - Image uuid for identifying previous base image.
* `previous_patch_image_ref` - Image uuid for identifying previous patch.
* `previous_patch_version` - Previous patch version applied to this node.
* `previous_version` - Previous version prior to upgrade.
* `se_upgrade_events` - Serviceenginegroup upgrade errors.
* `seg_status` - Detailed segroup status.
* `start_time` - Start time of upgrade operations.
* `state` - Current status of the upgrade operations.
* `system` - Flag is set only in the cluster node entry if the upgrade is initiated as a system-upgrade.
* `tasks_completed` - Upgrade tasks completed.
* `tenant_ref` - Tenant that this object belongs to.
* `total_tasks` - Total upgrade tasks.
* `upgrade_events` - Events performed for upgrade operations.
* `upgrade_ops` - Upgrade operations requested.
* `uuid` - Uuid identifier for the system such as cluster, se group and se.
* `version` - Current base image applied to this node.

