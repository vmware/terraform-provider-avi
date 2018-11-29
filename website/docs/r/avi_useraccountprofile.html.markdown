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
resource "UserAccountProfile" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `account_lock_timeout` - (Optional ) argument_description.
        * `credentials_timeout_threshold` - (Optional ) argument_description.
        * `max_concurrent_sessions` - (Optional ) argument_description.
        * `max_login_failure_count` - (Optional ) argument_description.
        * `max_password_history_count` - (Optional ) argument_description.
        * `name` - (Required) argument_description.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                            * `uuid` - argument_description.
