---
layout: "avi"
page_title: "Avi: avi_serviceenginepolicy"
sidebar_current: "docs-avi-resource-serviceenginepolicy"
description: |-
  Creates and manages Avi ServiceEnginePolicy.
---

# avi_serviceenginepolicy

The ServiceEnginePolicy resource allows the creation and management of Avi ServiceEnginePolicy

## Example Usage

```hcl
resource "ServiceEnginePolicy" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `name` - (Optional ) argument_description.
        * `nat_policy_ref` - (Optional ) argument_description.
        * `se_group_ref` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
            * `vrf_ref` - (Optional ) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                    * `uuid` - argument_description.
        
