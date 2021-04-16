<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_systemlimits"
sidebar_current: "docs-avi-datasource-systemlimits"
description: |-
  Get information of Avi SystemLimits.
---

# avi_systemlimits

This data source is used to to get avi_systemlimits objects.

## Example Usage

```hcl
data "avi_systemlimits" "foo_systemlimits" {
    uuid = "systemlimits-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SystemLimits by name.
* `uuid` - (Optional) Search SystemLimits by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `controller_limits` - System limits for the entire controller cluster. Field introduced in 20.1.1.
* `controller_sizes` - Possible controller sizes. Field introduced in 20.1.1.
* `serviceengine_limits` - System limits that apply to a serviceengine. Field introduced in 20.1.1.
* `uuid` - Uuid for the system limits object. Field introduced in 20.1.1.

