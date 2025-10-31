<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_passwordpolicy"
sidebar_current: "docs-avi-datasource-passwordpolicy"
description: |-
  Get information of Avi PasswordPolicy.
---

# avi_passwordpolicy

This data source is used to to get avi_passwordpolicy objects.

## Example Usage

```hcl
data "avi_passwordpolicy" "foo_passwordpolicy" {
    uuid = "passwordpolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search PasswordPolicy by name.
* `uuid` - (Optional) Search PasswordPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 32.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `lockout_evaluation_period` - Time window for evaluating failed attempts in seconds. Defaults to 900 seconds. Allowed values are 300-1800. Field introduced in 32.1.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.
* `lockout_max_auth_failures` - Number of failed attempts before account lockout. Defaults to 3. Allowed values are 0-5. Special values are 0- unlimited login attempts allowed. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `lockout_period` - Account lockout duration in seconds. Defaults to 900 seconds. Allowed values are 600-1800. Field introduced in 32.1.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_length` - Minimum password length. Defaults to 15 characters. Allowed values are 8-64. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_lowercase` - Minimum number of lowercase characters required. Allowed values are 0-10. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_numeric` - Minimum number of numeric characters required. Allowed values are 0-10. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_special` - Minimum number of special characters required. Allowed values are 0-10. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_uppercase` - Minimum number of uppercase characters required. Allowed values are 0-10. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - Name of the password policy configuration. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `password_expiration_days` - Password expiry period in days. Defaults to 365 days. Allowed values are 30-730. Field introduced in 32.1.1. Unit is days. Allowed with any value in enterprise, enterprise with cloud services edition.
* `password_history` - Number of previous passwords to remember. Defaults to 5. Allowed values are 1-10. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Tenant ref for the passwordpolicy. It is a reference to an object of type tenant. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Unique object identifier of the passwordpolicy. Field introduced in 32.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

