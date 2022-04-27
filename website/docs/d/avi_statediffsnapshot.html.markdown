<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_statediffsnapshot"
sidebar_current: "docs-avi-datasource-statediffsnapshot"
description: |-
  Get information of Avi StatediffSnapshot.
---

# avi_statediffsnapshot

This data source is used to to get avi_statediffsnapshot objects.

## Example Usage

```hcl
data "avi_statediffsnapshot" "foo_statediffsnapshot" {
    uuid = "statediffsnapshot-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search StatediffSnapshot by name.
* `uuid` - (Optional) Search StatediffSnapshot by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `gslb_name` - Name of gslb object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `gslb_uuid` - Reference to base gslb object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Name of statediff operation. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `pool_name` - Name of pool object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `pool_uuid` - Reference to base pool object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `post_snapshot` - Post-upgrade snapshot for vs. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `pre_snapshot` - Pre-upgrade snapshot for vs. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `se_group_name` - Name of seg object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `se_group_uuid` - Reference to base seg object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `se_name` - Name of seg object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `se_uuid` - Reference to base se object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `snapshot_type` - Type of snapshot eg. Vs_snapshot, se_snapshot etc. Enum options - FB_VS_SNAPSHOT, FB_SE_SNAPSHOT, FB_GSLB_SNAPSHOT, FB_POOL_SNAPSHOT. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `statediff_operation_ref` - Statediff operation uuid for identifying the operation. It is a reference to an object of type statediffoperation. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Unique identifier for statediff entry. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vs_name` - Name of vs object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vs_uuid` - Reference to base vs object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.

