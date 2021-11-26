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
page_title: "AVI: avi_statediffoperation"
sidebar_current: "docs-avi-datasource-statediffoperation"
description: |-
  Get information of Avi StatediffOperation.
---

# avi_statediffoperation

This data source is used to to get avi_statediffoperation objects.

## Example Usage

```hcl
data "avi_statediffoperation" "foo_statediffoperation" {
    uuid = "statediffoperation-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search StatediffOperation by name.
* `uuid` - (Optional) Search StatediffOperation by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `events` - Info for each statediff event. Field introduced in 21.1.3.
* `name` - Name of statediff operation. Field introduced in 21.1.3.
* `node_uuid` - Uuid of node for statediff operation entry. Field introduced in 21.1.3.
* `operation` - Type of statediff operation. Enum options - FB_UPGRADE, FB_ROLLBACK, FB_PATCH, FB_ROLLBACK_PATCH. Field introduced in 21.1.3.
* `phase` - Phase of statediff operation. Enum options - FB_PRE_SNAPSHOT, FB_POST_SNAPSHOT. Field introduced in 21.1.3.
* `status` - Status of statediff operation. Enum options - FB_INIT, FB_IN_PROGRESS, FB_COMPLETED, FB_FAILED, FB_COMPLETED_WITH_ERRORS. Field introduced in 21.1.3.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 21.1.3.
* `uuid` - Unique identifier for statediff entry. Field introduced in 21.1.3.

