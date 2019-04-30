---
layout: "avi"
page_title: "AVI: avi_ssopolicy"
sidebar_current: "docs-avi-datasource-ssopolicy"
description: |-
  Get information of Avi SSOPolicy.
---

# avi_ssopolicy

This data source is used to to get avi_ssopolicy objects.

## Example Usage

```hcl
data "SSOPolicy" "foo_SSOPolicy" {
    uuid = "SSOPolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SSOPolicy by name.
* `uuid` - (Optional) Search SSOPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `authentication_policy` - Authentication policy settings.
* `name` - Name of the sso policy.
* `tenant_ref` - Uuid of the tenant.
* `uuid` - Uuid of the sso policy.

