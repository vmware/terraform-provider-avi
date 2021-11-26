############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

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

* `access` - List of list.
* `default_tenant_ref` - It is a reference to an object of type tenant.
* `email` - Placeholder for description of property email of obj type user field type string  type str.
* `full_name` - Placeholder for description of property full_name of obj type user field type string  type str.
* `is_superuser` - Boolean flag to set is_superuser.
* `local` - Boolean flag to set local.
* `name` - Name of the object.
* `password` - Placeholder for description of property password of obj type user field type string  type str.
* `user_profile_ref` - It is a reference to an object of type useraccountprofile.
* `username` - Placeholder for description of property username of obj type user field type string  type str.
* `uuid` - Unique object identifier of the object.

