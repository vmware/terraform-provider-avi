<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_botmapping"
sidebar_current: "docs-avi-resource-botmapping"
description: |-
  Creates and manages Avi BotMapping.
---

# avi_botmapping

The BotMapping resource allows the creation and management of Avi BotMapping

## Example Usage

```hcl
resource "avi_botmapping" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of this mapping. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `mapping_rules` - (Optional) Rules for bot classification. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) The unique identifier of the tenant to which this mapping belongs. It is a reference to an object of type tenant. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  A unique identifier for this mapping. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

