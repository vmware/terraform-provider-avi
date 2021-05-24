<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_rmcloudopsproto"
sidebar_current: "docs-avi-resource-rmcloudopsproto"
description: |-
  Creates and manages Avi RmCloudOpsProto.
---

# avi_rmcloudopsproto

The RmCloudOpsProto resource allows the creation and management of Avi RmCloudOpsProto

## Example Usage

```hcl
resource "avi_rmcloudopsproto" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `last_queried_se_creation_limit` - (Optional) The most recent value of concurrent se creation limit from cloudconnectorstatus. Field introduced in 20.1.1.
* `name` - (Optional) Cloud name. Field introduced in 20.1.1.
* `pending_se_creation_count` - (Optional) Number of se creations in progress. Field introduced in 20.1.1.
* `pending_vnic_op_count` - (Optional) Number of vnic operations in progress (both add and delete). Field introduced in 20.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Cloud uuid. Field introduced in 20.1.1.

