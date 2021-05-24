<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_labelgroup"
sidebar_current: "docs-avi-datasource-labelgroup"
description: |-
  Get information of Avi LabelGroup.
---

# avi_labelgroup

This data source is used to to get avi_labelgroup objects.

## Example Usage

```hcl
data "avi_labelgroup" "foo_labelgroup" {
    uuid = "labelgroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search LabelGroup by name.
* `uuid` - (Optional) Search LabelGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `labels` - List of allowed or suggested labels for the label group. Field introduced in 20.1.5.
* `name` - Name of the label group. Field introduced in 20.1.5.
* `uuid` - Uuid of the label group. Field introduced in 20.1.5.

