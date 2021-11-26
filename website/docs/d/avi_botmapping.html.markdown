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
page_title: "AVI: avi_botmapping"
sidebar_current: "docs-avi-datasource-botmapping"
description: |-
  Get information of Avi BotMapping.
---

# avi_botmapping

This data source is used to to get avi_botmapping objects.

## Example Usage

```hcl
data "avi_botmapping" "foo_botmapping" {
    uuid = "botmapping-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search BotMapping by name.
* `uuid` - (Optional) Search BotMapping by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `mapping_rules` - Rules for bot classification. Field introduced in 21.1.1.
* `name` - The name of this mapping. Field introduced in 21.1.1.
* `tenant_ref` - The unique identifier of the tenant to which this mapping belongs. It is a reference to an object of type tenant. Field introduced in 21.1.1.
* `uuid` - A unique identifier for this mapping. Field introduced in 21.1.1.

