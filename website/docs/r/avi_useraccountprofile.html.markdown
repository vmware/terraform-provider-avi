<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_useraccountprofile"
sidebar_current: "docs-avi-resource-useraccountprofile"
description: |-
  Creates and manages Avi UserAccountProfile.
---

# avi_useraccountprofile

The UserAccountProfile resource allows the creation and management of Avi UserAccountProfile

## Example Usage

```hcl
resource "avi_useraccountprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `account_lock_timeout` - (Optional) Lock timeout period (in minutes). Default is 30 minutes. Unit is min.
* `credentials_timeout_threshold` - (Optional) The time period after which credentials expire. Default is 180 days. Unit is days.
* `max_concurrent_sessions` - (Optional) Maximum number of concurrent sessions allowed. There are unlimited sessions by default.
* `max_login_failure_count` - (Optional) Number of login attempts before lockout. Default is 3 attempts. Allowed values are 3-20. Special values are 0 - 'unlimited login attempts allowed.'.
* `max_password_history_count` - (Optional) Maximum number of passwords to be maintained in the password history. Default is 4 passwords.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

