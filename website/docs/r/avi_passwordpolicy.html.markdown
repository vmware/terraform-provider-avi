<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_passwordpolicy"
sidebar_current: "docs-avi-resource-passwordpolicy"
description: |-
  Creates and manages Avi PasswordPolicy.
---

# avi_passwordpolicy

The PasswordPolicy resource allows the creation and management of Avi PasswordPolicy

## Example Usage

```hcl
resource "avi_passwordpolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the password policy configuration. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 31.3.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `lockout_evaluation_period` - (Optional) Time window for evaluating failed attempts in seconds. Defaults to 900 seconds. Allowed values are 300-1800. Field introduced in 31.3.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.
* `lockout_max_auth_failures` - (Optional) Number of failed attempts before account lockout. Defaults to 3. Allowed values are 0-5. Special values are 0- unlimited login attempts allowed. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `lockout_period` - (Optional) Account lockout duration in seconds. Defaults to 900 seconds. Allowed values are 600-1800. Field introduced in 31.3.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_length` - (Optional) Minimum password length. Defaults to 15 characters. Allowed values are 8-64. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_lowercase` - (Optional) Minimum number of lowercase characters required. Allowed values are 0-10. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_numeric` - (Optional) Minimum number of numeric characters required. Allowed values are 0-10. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_special` - (Optional) Minimum number of special characters required. Allowed values are 0-10. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `min_uppercase` - (Optional) Minimum number of uppercase characters required. Allowed values are 0-10. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `password_expiration_days` - (Optional) Password expiry period in days. Defaults to 365 days. Allowed values are 30-730. Field introduced in 31.3.1. Unit is days. Allowed with any value in enterprise, enterprise with cloud services edition.
* `password_history` - (Optional) Number of previous passwords to remember. Defaults to 5. Allowed values are 1-10. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant ref for the passwordpolicy. It is a reference to an object of type tenant. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the passwordpolicy. Field introduced in 31.3.1. Allowed with any value in enterprise, enterprise with cloud services edition.

