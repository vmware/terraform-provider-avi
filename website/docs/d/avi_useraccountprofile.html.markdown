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

* `account_lock_timeout` - Lock timeout period (in minutes). Default is 30 minutes. Unit is min. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `credentials_timeout_threshold` - The time period after which credentials expire. Default is 180 days. Unit is days. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `login_failure_count_expiry_window` - The configurable time window beyond which we need to pop all the login failure timestamps from the login_failure_timestamps. Special values are 0 - do not reset login_failure_counts on the basis of time. Field introduced in 22.1.1. Unit is min. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `max_concurrent_sessions` - Maximum number of concurrent sessions allowed. There are unlimited sessions by default. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `max_login_failure_count` - Number of login attempts before lockout. Default is 3 attempts. Allowed values are 3-20. Special values are 0- unlimited login attempts allowed. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `max_password_history_count` - Maximum number of passwords to be maintained in the password history. Default is 4 passwords. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

