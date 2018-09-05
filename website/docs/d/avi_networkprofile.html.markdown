---
layout: "avi"
page_title: "AVI: avi_networkprofile"
sidebar_current: "docs-avi-datasource-networkprofile"
description: |-
Get information of Avi NetworkProfile.
---

# avi_networkprofile

This data source is used to to get avi_networkprofile objects.

## Example Usage

```hcl
data "NetworkProfile" "foo_NetworkProfile" {
    uuid = "NetworkProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search NetworkProfile by name.
* `uuid` - (Optional) Search NetworkProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `connection_mirror` - When enabled, avi mirrors all tcp fastpath connections to standby.applicable only in legacy ha mode.
* `description` - General description.
* `name` - The name of the network profile.
* `profile` - General description.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the network profile.

