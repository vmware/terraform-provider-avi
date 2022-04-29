<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_user"
sidebar_current: "docs-avi-datasource-user"
description: |-
  Get information of Avi User.
---

# avi_user

This data source is used to to get avi_user objects.

## Example Usage

```hcl
data "avi_user" "foo_user" {
    uuid = "user-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search User by name.
* `uuid` - (Optional) Search User by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `access` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `default_tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `email` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `full_name` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `is_superuser` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `local` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `password` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `user_profile_ref` - It is a reference to an object of type useraccountprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `username` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

