<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_statediffoperation"
sidebar_current: "docs-avi-resource-statediffoperation"
description: |-
  Creates and manages Avi StatediffOperation.
---

# avi_statediffoperation

The StatediffOperation resource allows the creation and management of Avi StatediffOperation

## Example Usage

```hcl
resource "avi_statediffoperation" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.3.
* `events` - (Optional) Info for each statediff event. Field introduced in 21.1.3.
* `name` - (Optional) Name of statediff operation. Field introduced in 21.1.3.
* `node_uuid` - (Optional) Uuid of node for statediff operation entry. Field introduced in 21.1.3.
* `operation` - (Optional) Type of statediff operation. Enum options - FB_UPGRADE, FB_ROLLBACK, FB_PATCH, FB_ROLLBACK_PATCH. Field introduced in 21.1.3.
* `phase` - (Optional) Phase of statediff operation. Enum options - FB_PRE_SNAPSHOT, FB_POST_SNAPSHOT. Field introduced in 21.1.3.
* `status` - (Optional) Status of statediff operation. Enum options - FB_INIT, FB_IN_PROGRESS, FB_COMPLETED, FB_FAILED, FB_COMPLETED_WITH_ERRORS. Field introduced in 21.1.3.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 21.1.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique identifier for statediff entry. Field introduced in 21.1.3.

