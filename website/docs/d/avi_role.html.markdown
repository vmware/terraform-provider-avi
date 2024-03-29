<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
data "avi_role" "foo_role" {
    uuid = "role-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Role by name.
* `uuid` - (Optional) Search Role by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `allow_unlabelled_access` - Allow access to unlabelled objects. Field introduced in 20.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `filters` - Filters for granular object access control based on object labels. Multiple filters are merged using the and operator. If empty, all objects according to the privileges will be accessible to the user. Field introduced in 20.1.3. Maximum of 4 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `privileges` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

