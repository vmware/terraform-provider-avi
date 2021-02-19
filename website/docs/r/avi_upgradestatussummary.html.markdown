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

* `enable_patch_rollback` - (Optional) Check if the patch rollback is possible on this node. Field introduced in 18.2.6.
* `enable_rollback` - (Optional) Check if the rollback is possible on this node. Field introduced in 18.2.6.
* `end_time` - (Optional) End time of upgrade operations. Field introduced in 18.2.6.
* `image_ref` - (Optional) Image uuid for identifying the current base image. It is a reference to an object of type image. Field introduced in 18.2.6.
* `name` - (Optional) Name of the system such as cluster name, se group name and se name. Field introduced in 18.2.6.
* `node_type` - (Optional) Type of the system such as controller_cluster, se_group or se. Enum options - NODE_CONTROLLER_CLUSTER, NODE_SE_GROUP, NODE_SE_TYPE. Field introduced in 18.2.6.
* `obj_cloud_ref` - (Optional) Cloud that this object belongs to. It is a reference to an object of type cloud. Field introduced in 18.2.6.
* `patch_image_ref` - (Optional) Image uuid for identifying the current patch. It is a reference to an object of type image. Field introduced in 18.2.6.
* `start_time` - (Optional) Start time of upgrade operations. Field introduced in 18.2.6.
* `state` - (Optional) Current status of the upgrade operations. Field introduced in 18.2.6.
* `tasks_completed` - (Optional) Upgrade tasks completed. Field introduced in 18.2.6.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.6.
* `total_tasks` - (Optional) Total upgrade tasks. Field introduced in 18.2.6.
* `upgrade_ops` - (Optional) Upgrade operations requested. Enum options - UPGRADE, PATCH, ROLLBACK, ROLLBACKPATCH, SEGROUP_RESUME. Field introduced in 18.2.6.
* `version` - (Optional) Current base image applied to this node. Field introduced in 18.2.6.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the system such as cluster, se group and se. Field introduced in 18.2.6.

