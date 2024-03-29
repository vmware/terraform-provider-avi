<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_botconfigconsolidator"
sidebar_current: "docs-avi-resource-botconfigconsolidator"
description: |-
  Creates and manages Avi BotConfigConsolidator.
---

# avi_botconfigconsolidator

The BotConfigConsolidator resource allows the creation and management of Avi BotConfigConsolidator

## Example Usage

```hcl
resource "avi_botconfigconsolidator" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of this consolidator. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `description` - (Optional) Human-readable description of this consolidator. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `script` - (Optional) Script that consolidates results from all bot decision components. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) The unique identifier of the tenant to which this consolidator belongs. It is a reference to an object of type tenant. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  A unique identifier for this consolidator. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

