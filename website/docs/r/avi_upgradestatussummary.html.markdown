############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_upgradestatussummary"
sidebar_current: "docs-avi-resource-upgradestatussummary"
description: |-
  Creates and manages Avi UpgradeStatusSummary.
---

# avi_upgradestatussummary

The UpgradeStatusSummary resource allows the creation and management of Avi UpgradeStatusSummary

## Example Usage

```hcl
resource "avi_upgradestatussummary" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `enable_patch_rollback` - (Optional) Check if the patch rollback is possible on this node.
* `enable_rollback` - (Optional) Check if the rollback is possible on this node.
* `end_time` - (Optional) End time of upgrade operations.
* `image_ref` - (Optional) Image uuid for identifying the current base image.
* `name` - (Optional) Name of the system such as cluster name, se group name and se name.
* `node_type` - (Optional) Type of the system such as controller_cluster, se_group or se.
* `obj_cloud_ref` - (Optional) Cloud that this object belongs to.
* `patch_image_ref` - (Optional) Image uuid for identifying the current patch.
* `start_time` - (Optional) Start time of upgrade operations.
* `state` - (Optional) Current status of the upgrade operations.
* `tasks_completed` - (Optional) Upgrade tasks completed.
* `tenant_ref` - (Optional) Tenant that this object belongs to.
* `total_tasks` - (Optional) Total upgrade tasks.
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

