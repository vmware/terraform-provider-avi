---
layout: "avi"
page_title: "Avi: avi_upgradestatusinfo"
sidebar_current: "docs-avi-resource-upgradestatusinfo"
description: |-
  Creates and manages Avi UpgradeStatusInfo.
---

# avi_upgradestatusinfo

The UpgradeStatusInfo resource allows the creation and management of Avi UpgradeStatusInfo

## Example Usage

```hcl
resource "avi_upgradestatusinfo" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `duration` - (Optional) Duration of upgrade operations in seconds.
* `enable_patch_rollback` - (Optional) Check if the patch rollback is possible on this node.
* `enable_rollback` - (Optional) Check if the rollback is possible on this node.
* `end_time` - (Optional) End time of upgrade operations.
* `enqueue_time` - (Optional) Enqueue time of upgrade operations.
* `image_ref` - (Optional) Image uuid for identifying the current base image.
* `name` - (Optional) Name of the system such as cluster name, se group name and se name.
* `node_type` - (Optional) Type of the system such as controller_cluster, se_group or se.
* `obj_cloud_ref` - (Optional) Cloud that this object belongs to.
* `params` - (Optional) Parameters associated with the current upgrade ops.
* `patch_image_ref` - (Optional) Image uuid for identifying the current patch.
* `patch_version` - (Optional) Current patch version applied to this node.
* `previous_image_ref` - (Optional) Image uuid for identifying previous base image.
* `previous_patch_image_ref` - (Optional) Image uuid for identifying previous patch.
* `previous_patch_version` - (Optional) Previous patch version applied to this node.
* `previous_version` - (Optional) Previous version prior to upgrade.
* `se_upgrade_events` - (Optional) Serviceenginegroup upgrade errors.
* `seg_status` - (Optional) Detailed segroup status.
* `start_time` - (Optional) Start time of upgrade operations.
* `state` - (Optional) Current status of the upgrade operations.
* `system` - (Optional) Flag is set only in the cluster node entry if the upgrade is initiated as a system-upgrade.
* `tasks_completed` - (Optional) Upgrade tasks completed.
* `tenant_ref` - (Optional) Tenant that this object belongs to.
* `total_tasks` - (Optional) Total upgrade tasks.
* `upgrade_events` - (Optional) Events performed for upgrade operations.
* `upgrade_ops` - (Optional) Upgrade operations requested.
* `version` - (Optional) Current base image applied to this node.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the system such as cluster, se group and se.

