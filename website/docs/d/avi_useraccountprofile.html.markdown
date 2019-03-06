
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
###

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
data "UserAccountProfile" "foo_UserAccountProfile" {
    uuid = "UserAccountProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search UserAccountProfile by name.
* `uuid` - (Optional) Search UserAccountProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `account_lock_timeout` - Lock timeout period (in minutes).
* `credentials_timeout_threshold` - The time period after which credentials expire.
* `max_concurrent_sessions` - Maximum number of concurrent sessions allowed.
* `max_login_failure_count` - Number of login attempts before lockout.
* `max_password_history_count` - Maximum number of passwords to be maintained in the password history.
* `name` - General description.
* `uuid` - General description.

