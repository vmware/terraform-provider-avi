---
layout: "avi"
page_title: "Avi: avi_hardwaresecuritymodulegroup"
sidebar_current: "docs-avi-resource-hardwaresecuritymodulegroup"
description: |-
  Creates and manages Avi HardwareSecurityModuleGroup.
---

# avi_hardwaresecuritymodulegroup

The HardwareSecurityModuleGroup resource allows the creation and management of Avi HardwareSecurityModuleGroup

## Example Usage

```hcl
resource "HardwareSecurityModuleGroup" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `hsm` - (Required) argument_description.
        * `name` - (Required) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                * `uuid` - argument_description.
    
