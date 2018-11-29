---
layout: "avi"
page_title: "Avi: avi_pkiprofile"
sidebar_current: "docs-avi-resource-pkiprofile"
description: |-
  Creates and manages Avi PKIProfile.
---

# avi_pkiprofile

The PKIProfile resource allows the creation and management of Avi PKIProfile

## Example Usage

```hcl
resource "PKIProfile" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `ca_certs` - (Optional ) argument_description.
        * `created_by` - (Optional ) argument_description.
        * `crl_check` - (Optional ) argument_description.
        * `crls` - (Optional ) argument_description.
        * `ignore_peer_chain` - (Optional ) argument_description.
        * `is_federated` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
            * `validate_only_leaf_crl` - (Optional ) argument_description.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                    * `uuid` - argument_description.
