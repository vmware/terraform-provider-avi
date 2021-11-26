############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `enable_patch_rollback` - Check if the patch rollback is possible on this node. Field introduced in 18.2.6.
* `enable_rollback` - Check if the rollback is possible on this node. Field introduced in 18.2.6.
* `end_time` - End time of upgrade operations. Field introduced in 18.2.6.
* `image_ref` - Image uuid for identifying the current base image. It is a reference to an object of type image. Field introduced in 18.2.6.
* `name` - Name of the system such as cluster name, se group name and se name. Field introduced in 18.2.6.
* `node_type` - Type of the system such as controller_cluster, se_group or se. Enum options - NODE_CONTROLLER_CLUSTER, NODE_SE_GROUP, NODE_SE_TYPE. Field introduced in 18.2.6.
* `obj_cloud_ref` - Cloud that this object belongs to. It is a reference to an object of type cloud. Field introduced in 18.2.6.
* `patch_image_ref` - Image uuid for identifying the current patch. It is a reference to an object of type image. Field introduced in 18.2.6.
* `start_time` - Start time of upgrade operations. Field introduced in 18.2.6.
* `state` - Current status of the upgrade operations. Field introduced in 18.2.6.
* `tasks_completed` - Upgrade tasks completed. Field introduced in 18.2.6.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.6.
* `total_tasks` - Total upgrade tasks. Field introduced in 18.2.6.
* `upgrade_ops` - Upgrade operations requested. Enum options - UPGRADE, PATCH, ROLLBACK, ROLLBACKPATCH, SEGROUP_RESUME. Field introduced in 18.2.6.
* `uuid` - Uuid identifier for the system such as cluster, se group and se. Field introduced in 18.2.6.
* `version` - Current base image applied to this node. Field introduced in 18.2.6.

