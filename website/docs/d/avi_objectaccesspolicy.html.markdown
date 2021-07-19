<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_objectaccesspolicy"
sidebar_current: "docs-avi-datasource-objectaccesspolicy"
description: |-
  Get information of Avi ObjectAccessPolicy.
---

# avi_objectaccesspolicy

This data source is used to to get avi_objectaccesspolicy objects.

## Example Usage

```hcl
data "avi_objectaccesspolicy" "foo_objectaccesspolicy" {
    uuid = "objectaccesspolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ObjectAccessPolicy by name.
* `uuid` - (Optional) Search ObjectAccessPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of the object access policy. Field introduced in 18.2.7.
* `rules` - Rules which grant access to specific objects. Field introduced in 18.2.7.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.7.
* `uuid` - Uuid of the object access policy. Field introduced in 18.2.7.

