---
layout: "avi"
page_title: "Avi: avi_wafpolicypsmgroup"
sidebar_current: "docs-avi-resource-wafpolicypsmgroup"
description: |-
  Creates and manages Avi WafPolicyPSMGroup.
---

# avi_wafpolicypsmgroup

The WafPolicyPSMGroup resource allows the creation and management of Avi WafPolicyPSMGroup

## Example Usage

```hcl
resource "WafPolicyPSMGroup" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `description` - (Optional ) argument_description.
        * `enable` - (Optional ) argument_description.
        * `hit_action` - (Optional ) argument_description.
        * `is_learning_group` - (Optional ) argument_description.
        * `locations` - (Optional ) argument_description.
        * `miss_action` - (Optional ) argument_description.
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
    
