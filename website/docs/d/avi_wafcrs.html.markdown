---
layout: "avi"
page_title: "AVI: avi_wafcrs"
sidebar_current: "docs-avi-datasource-wafcrs"
description: |-
  Get information of Avi WafCRS.
---

# avi_wafcrs

This data source is used to to get avi_wafcrs objects.

## Example Usage

```hcl
data "avi_wafcrs" "foo_wafcrs" {
    uuid = "wafcrs-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafCRS by name.
* `uuid` - (Optional) Search WafCRS by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - A short description of this ruleset.
* `groups` - Waf rules are sorted in groups based on their characterization.
* `integrity` - Integrity protection value.
* `name` - The name of this ruleset object.
* `release_date` - The release date of this version in rfc 3339 / iso 8601 format.
* `tenant_ref` - Tenant that this object belongs to.
* `uuid` - Field introduced in 18.1.1.
* `version` - The version of this ruleset object.

