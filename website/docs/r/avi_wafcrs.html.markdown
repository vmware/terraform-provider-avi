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

* `description` - (Optional) A short description of this ruleset.
* `groups` - (Optional) Waf rules are sorted in groups based on their characterization.
* `integrity` - (Optional) Integrity protection value.
* `name` - (Optional) The name of this ruleset object.
* `release_date` - (Optional) The release date of this version in rfc 3339 / iso 8601 format.
* `tenant_ref` - (Optional) Tenant that this object belongs to.
* `version` - (Optional) The version of this ruleset object.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 18.1.1.

