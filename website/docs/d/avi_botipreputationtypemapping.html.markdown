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
page_title: "AVI: avi_botipreputationtypemapping"
sidebar_current: "docs-avi-datasource-botipreputationtypemapping"
description: |-
  Get information of Avi BotIPReputationTypeMapping.
---

# avi_botipreputationtypemapping

This data source is used to to get avi_botipreputationtypemapping objects.

## Example Usage

```hcl
data "avi_botipreputationtypemapping" "foo_botipreputationtypemapping" {
    uuid = "botipreputationtypemapping-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search BotIPReputationTypeMapping by name.
* `uuid` - (Optional) Search BotIPReputationTypeMapping by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `ip_reputation_mappings` - Map every ipreputationtype to a bot type (can be unknown). Field introduced in 21.1.1.
* `name` - The name of this mapping. Field introduced in 21.1.1.
* `tenant_ref` - The unique identifier of the tenant to which this mapping belongs. It is a reference to an object of type tenant. Field introduced in 21.1.1.
* `uuid` - A unique identifier of this mapping. Field introduced in 21.1.1.

