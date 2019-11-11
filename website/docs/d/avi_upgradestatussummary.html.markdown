---
layout: "avi"
page_title: "AVI: avi_upgradestatussummary"
sidebar_current: "docs-avi-datasource-upgradestatussummary"
description: |-
  Get information of Avi UpgradeStatusSummary.
---

# avi_upgradestatussummary

This data source is used to to get avi_upgradestatussummary objects.

## Example Usage

```hcl
data "avi_upgradestatussummary" "foo_upgradestatussummary" {
    uuid = "upgradestatussummary-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search UpgradeStatusSummary by name.
* `uuid` - (Optional) Search UpgradeStatusSummary by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `enable_patch_rollback` - Check if the patch rollback is possible on this node.
* `enable_rollback` - Check if the rollback is possible on this node.
* `end_time` - End time of upgrade operations.
* `image_ref` - Image uuid for identifying the current base image.
* `name` - Name of the system such as cluster name, se group name and se name.
* `node_type` - Type of the system such as controller_cluster, se_group or se.
* `obj_cloud_ref` - Cloud that this object belongs to.
* `patch_image_ref` - Image uuid for identifying the current patch.
* `start_time` - Start time of upgrade operations.
* `state` - Current status of the upgrade operations.
* `tasks_completed` - Upgrade tasks completed.
* `tenant_ref` - Tenant that this object belongs to.
* `total_tasks` - Total upgrade tasks.
* `upgrade_ops` - Upgrade operations requested.
* `uuid` - Uuid identifier for the system such as cluster, se group and se.
* `version` - Current base image applied to this node.

