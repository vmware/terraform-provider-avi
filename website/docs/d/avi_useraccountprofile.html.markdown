<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_useraccountprofile"
sidebar_current: "docs-avi-datasource-useraccountprofile"
description: |-
  Get information of Avi UserAccountProfile.
---

# avi_useraccountprofile

This data source is used to to get avi_useraccountprofile objects.

## Example Usage

```hcl
data "avi_useraccountprofile" "foo_useraccountprofile" {
    uuid = "useraccountprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search UserAccountProfile by name.
* `uuid` - (Optional) Search UserAccountProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `account_lock_timeout` - Lock timeout period (in minutes). Default is 30 minutes.
* `credentials_timeout_threshold` - The time period after which credentials expire. Default is 180 days.
* `max_concurrent_sessions` - Maximum number of concurrent sessions allowed. There are unlimited sessions by default.
* `max_login_failure_count` - Number of login attempts before lockout. Default is 3 attempts.
* `max_password_history_count` - Maximum number of passwords to be maintained in the password history. Default is 4 passwords.
* `name` - Name of the object.
* `uuid` - Unique object identifier of the object.

