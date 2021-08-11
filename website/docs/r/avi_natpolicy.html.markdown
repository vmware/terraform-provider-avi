<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_natpolicy"
sidebar_current: "docs-avi-resource-natpolicy"
description: |-
  Creates and manages Avi NatPolicy.
---

# avi_natpolicy

The NatPolicy resource allows the creation and management of Avi NatPolicy

## Example Usage

```hcl
resource "avi_natpolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `created_by` - (Optional) Creator name. Field introduced in 18.2.3.
* `description` - (Optional) Field introduced in 18.2.3.
* `name` - (Optional) Name of the nat policy. Field introduced in 18.2.3.
* `rules` - (Optional) Nat policy rules. Field introduced in 18.2.3.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 18.2.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the nat policy. Field introduced in 18.2.3.

