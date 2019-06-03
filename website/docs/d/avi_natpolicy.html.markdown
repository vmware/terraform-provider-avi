---
layout: "avi"
page_title: "AVI: avi_natpolicy"
sidebar_current: "docs-avi-datasource-natpolicy"
description: |-
  Get information of Avi NatPolicy.
---

# avi_natpolicy

This data source is used to to get avi_natpolicy objects.

## Example Usage

```hcl
data "avi_natpolicy" "foo_natpolicy" {
    uuid = "natpolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search NatPolicy by name.
* `uuid` - (Optional) Search NatPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `created_by` - Creator name.
* `description` - Field introduced in 18.2.3.
* `name` - Name of the nat policy.
* `rules` - Nat policy rules.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the nat policy.

