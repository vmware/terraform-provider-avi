---
layout: "avi"
page_title: "Avi: avi_l4policyset"
sidebar_current: "docs-avi-resource-l4policyset"
description: |-
  Creates and manages Avi L4PolicySet.
---

# avi_l4policyset

The L4PolicySet resource allows the creation and management of Avi L4PolicySet

## Example Usage

```hcl
resource "L4PolicySet" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `created_by` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `is_internal_policy` - (Optional ) argument_description.
        * `l4_connection_policy` - (Optional ) argument_description.
        * `name` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                            * `uuid` - argument_description.
    
