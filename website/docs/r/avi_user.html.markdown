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

* `name` - (Required) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `access` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `anonymous_user` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `date_joined` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `default_tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `email` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `full_name` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `is_active` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `is_internal_user` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `is_staff` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `is_superuser` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `local` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `logged_in` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `password` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `passwordless` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `recovery_token` - (Optional) Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `token_expiration_date` - (Optional) Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `ui_property` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uid` - (Optional) Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `unix_crypt_password` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `user_profile_ref` - (Optional) It is a reference to an object of type useraccountprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `username` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

