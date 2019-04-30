---
layout: "avi"
page_title: "Avi: avi_vrfcontext"
sidebar_current: "docs-avi-resource-vrfcontext"
description: |-
  Creates and manages Avi VrfContext.
---

# avi_vrfcontext

The VrfContext resource allows the creation and management of Avi VrfContext

## Example Usage

```hcl
resource "VrfContext" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `bgp_profile` - (Optional ) argument_description.
        * `cloud_ref` - (Optional ) argument_description.
        * `debugvrfcontext` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `gateway_mon` - (Optional ) argument_description.
        * `internal_gateway_monitor` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `static_routes` - (Optional ) argument_description.
        * `system_default` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                            * `uuid` - argument_description.
    
