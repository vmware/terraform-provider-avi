<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_wafcrs"
sidebar_current: "docs-avi-resource-wafcrs"
description: |-
  Creates and manages Avi WafCRS.
---

# avi_wafcrs

The WafCRS resource allows the creation and management of Avi WafCRS

## Example Usage

```hcl
resource "avi_wafcrs" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `description` - (Required) A short description of this ruleset. Field introduced in 18.1.1.
* `integrity` - (Required) Integrity protection value. Field introduced in 18.2.1.
* `name` - (Required) The name of this ruleset object. Field introduced in 18.2.1.
* `release_date` - (Required) The release date of this version in rfc 3339 / iso 8601 format. Field introduced in 18.1.1.
* `version` - (Required) The version of this ruleset object. Field introduced in 18.1.1.
* `groups` - (Optional) Waf rules are sorted in groups based on their characterization. Field introduced in 18.1.1.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 18.1.1.

