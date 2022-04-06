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

* `name` - (Required) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `account_lock_timeout` - (Optional) Lock timeout period (in minutes). Default is 30 minutes. Unit is min. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `credentials_timeout_threshold` - (Optional) The time period after which credentials expire. Default is 180 days. Unit is days. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `login_failure_count_expiry_window` - (Optional) The configurable time window beyond which we need to pop all the login failure timestamps from the login_failure_timestamps. Special values are 0 - do not reset login_failure_counts on the basis of time. Field introduced in 22.1.1. Unit is min. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `max_concurrent_sessions` - (Optional) Maximum number of concurrent sessions allowed. There are unlimited sessions by default. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `max_login_failure_count` - (Optional) Number of login attempts before lockout. Default is 3 attempts. Allowed values are 3-20. Special values are 0- unlimited login attempts allowed. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `max_password_history_count` - (Optional) Maximum number of passwords to be maintained in the password history. Default is 4 passwords. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

