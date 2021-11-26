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
page_title: "AVI: avi_rmcloudopsproto"
sidebar_current: "docs-avi-datasource-rmcloudopsproto"
description: |-
  Get information of Avi RmCloudOpsProto.
---

# avi_rmcloudopsproto

This data source is used to to get avi_rmcloudopsproto objects.

## Example Usage

```hcl
data "avi_rmcloudopsproto" "foo_rmcloudopsproto" {
    uuid = "rmcloudopsproto-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search RmCloudOpsProto by name.
* `uuid` - (Optional) Search RmCloudOpsProto by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `last_queried_se_creation_limit` - The most recent value of concurrent se creation limit from cloudconnectorstatus. Field introduced in 20.1.1.
* `name` - Cloud name. Field introduced in 20.1.1.
* `pending_se_creation_count` - Number of se creations in progress. Field introduced in 20.1.1.
* `pending_vnic_op_count` - Number of vnic operations in progress (both add and delete). Field introduced in 20.1.1.
* `uuid` - Cloud uuid. Field introduced in 20.1.1.

