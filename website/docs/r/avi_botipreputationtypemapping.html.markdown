<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_botipreputationtypemapping"
sidebar_current: "docs-avi-resource-botipreputationtypemapping"
description: |-
  Creates and manages Avi BotIPReputationTypeMapping.
---

# avi_botipreputationtypemapping

The BotIPReputationTypeMapping resource allows the creation and management of Avi BotIPReputationTypeMapping

## Example Usage

```hcl
resource "avi_botipreputationtypemapping" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of this mapping. Field introduced in 21.1.1.
* `ip_reputation_mappings` - (Optional) Map every ipreputationtype to a bot type (can be unknown). Field introduced in 21.1.1.
* `tenant_ref` - (Optional) The unique identifier of the tenant to which this mapping belongs. It is a reference to an object of type tenant. Field introduced in 21.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  A unique identifier of this mapping. Field introduced in 21.1.1.

