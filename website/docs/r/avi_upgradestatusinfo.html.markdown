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

* `duration` - (Optional) Duration of upgrade operation in seconds.
* `enable_patch_rollback` - (Optional) Check if the patch rollback is possible on this node.
* `enable_rollback` - (Optional) Check if the rollback is possible on this node.
* `end_time` - (Optional) End time of upgrade operation.
* `enqueue_time` - (Optional) Enqueue time of upgrade operation.
* `image_ref` - (Optional) Image uuid for identifying the current base image.
* `name` - (Optional) Name of the system such as cluster name, se group name and se name.
* `node_type` - (Optional) Type of the system such as controller_cluster, se_group or se.
* `obj_cloud_ref` - (Optional) Cloud that this object belongs to.
* `params` - (Optional) Parameters associated with the upgrade operation.
* `patch_image_ref` - (Optional) Image uuid for identifying the current patch.example  base-image is 18.2.6 and a patch 6p1 is applied, then this field will indicate the 6p1 value.
* `patch_list` - (Optional) List of patches applied to this node.
* `patch_version` - (Optional) Current patch version applied to this node.
* `previous_image_ref` - (Optional) Image uuid for identifying previous base image.example  base-image was 18.2.5 and an upgrade was done to 18.2.6, then this field will indicate the 18.2.5 value.
* `previous_patch_image_ref` - (Optional) Image uuid for identifying previous patch.example  base-image was 18.2.6 with a patch 6p1.
* `previous_patch_list` - (Optional) List of patches applied to this node on previous major version.
* `previous_patch_version` - (Optional) Previous patch version applied to this node.example  base-image was 18.2.6 with a patch 6p1.
* `previous_version` - (Optional) Previous version prior to upgrade.example  base-image was 18.2.5 and an upgrade was done to 18.2.6, then this field will indicate the 18.2.5 value.
* `progress` - (Optional) Upgrade operations progress which holds value between 0-100.
* `se_upgrade_events` - (Optional) Serviceenginegroup upgrade errors.
* `seg_status` - (Optional) Detailed segroup status.
* `start_time` - (Optional) Start time of upgrade operation.
* `state` - (Optional) Current status of the upgrade operation.
* `system` - (Optional) Flag is set only in the cluster if the upgrade is initiated as a system-upgrade.
* `tasks_completed` - (Optional) Completed set of tasks in the upgrade operation.
* `tenant_ref` - (Optional) Tenant that this object belongs to.
* `total_tasks` - (Optional) Total number of tasks in the upgrade operation.
* `upgrade_events` - (Optional) Events performed for upgrade operation.
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

