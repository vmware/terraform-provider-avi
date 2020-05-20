---
layout: "avi"
page_title: "Avi: avi_useraccount"
sidebar_current: "docs-avi-resource-useraccount"
description: |-
  Resource to reset user password.
---

# avi_useraccount

The Useraccount resource allows the password update of a user and setting up admin password in the bootstrap.

## Example Usage

```hcl
resource "avi_useraccount" "foo" {
    username     = var.avi_username	
    old_password = var.avi_current_password	
    password     = var.avi_new_password
}
```

## Argument Reference

The following arguments are supported:

    * `username` - (Required) User name of a user who's password needs to be updated.
    * `old_password` - (Required) Old/currant password of the user.
    * `password` - (Required) New paswword for the given user.
    * `local` - (Optional) Is local user.
    * `name` - (Optional) To set name for the user account.
    * `full_name` - (Optional) To set full name for the user account.
    * `email` - (Optional) To set email for the useraccount.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `update` - (Defaults to 40 mins) Used when updating the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

    * `uuid` - argument_description.
    