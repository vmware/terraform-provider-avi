<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_statediffsnapshot"
sidebar_current: "docs-avi-resource-statediffsnapshot"
description: |-
  Creates and manages Avi StatediffSnapshot.
---

# avi_statediffsnapshot

The StatediffSnapshot resource allows the creation and management of Avi StatediffSnapshot

## Example Usage

```hcl
resource "avi_statediffsnapshot" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `gslb_name` - (Optional) Name of gslb object. Field introduced in 21.1.3.
* `gslb_uuid` - (Optional) Reference to base gslb object. Field introduced in 21.1.3.
* `name` - (Optional) Name of statediff operation. Field introduced in 21.1.3.
* `pool_name` - (Optional) Name of pool object. Field introduced in 21.1.3.
* `pool_uuid` - (Optional) Reference to base pool object. Field introduced in 21.1.3.
* `post_snapshot` - (Optional) Post-upgrade snapshot for vs. Field introduced in 21.1.3.
* `pre_snapshot` - (Optional) Pre-upgrade snapshot for vs. Field introduced in 21.1.3.
* `se_group_name` - (Optional) Name of seg object. Field introduced in 21.1.3.
* `se_group_uuid` - (Optional) Reference to base seg object. Field introduced in 21.1.3.
* `se_name` - (Optional) Name of seg object. Field introduced in 21.1.3.
* `se_uuid` - (Optional) Reference to base se object. Field introduced in 21.1.3.
* `snapshot_type` - (Optional) Type of snapshot eg. Vs_snapshot, se_snapshot etc. Enum options - FB_VS_SNAPSHOT, FB_SE_SNAPSHOT, FB_GSLB_SNAPSHOT, FB_POOL_SNAPSHOT. Field introduced in 21.1.3.
* `statediff_operation_ref` - (Optional) Statediff operation uuid for identifying the operation. It is a reference to an object of type statediffoperation. Field introduced in 21.1.3.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 21.1.3.
* `vs_name` - (Optional) Name of vs object. Field introduced in 21.1.3.
* `vs_uuid` - (Optional) Reference to base vs object. Field introduced in 21.1.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique identifier for statediff entry. Field introduced in 21.1.3.

