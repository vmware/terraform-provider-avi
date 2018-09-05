---
layout: "avi"
page_title: "AVI: avi_role"
sidebar_current: "docs-avi-datasource-role"
description: |-
Get information of Avi Role.
---

# avi_role

This data source is used to to get avi_role objects.

## Example Usage

```hcl
data "Role" "foo_Role" {
    uuid = "Role-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Role by name.
* `uuid` - (Optional) Search Role by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - General description.
* `privileges` - General description.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.

