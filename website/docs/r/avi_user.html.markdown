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
page_title: "Avi: avi_user"
sidebar_current: "docs-avi-resource-user"
description: |-
  Creates and manages Avi User.
---

# avi_user

The User resource allows the creation and management of Avi User

## Example Usage

```hcl
resource "avi_user" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `access` - (Optional) List of list.
* `default_tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `email` - (Optional) Placeholder for description of property email of obj type user field type string  type str.
* `full_name` - (Optional) Placeholder for description of property full_name of obj type user field type string  type str.
* `is_superuser` - (Optional) Boolean flag to set is_superuser.
* `local` - (Optional) Boolean flag to set local.
* `password` - (Optional) Placeholder for description of property password of obj type user field type string  type str.
* `user_profile_ref` - (Optional) It is a reference to an object of type useraccountprofile.
* `username` - (Optional) Placeholder for description of property username of obj type user field type string  type str.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

