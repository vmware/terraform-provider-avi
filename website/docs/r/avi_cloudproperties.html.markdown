---
layout: "avi"
page_title: "Avi: avi_cloudproperties"
sidebar_current: "docs-avi-resource-cloudproperties"
description: |-
  Creates and manages Avi CloudProperties.
---

# avi_cloudproperties

The CloudProperties resource allows the creation and management of Avi CloudProperties

## Example Usage

```hcl
resource "CloudProperties" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `cc_props` - (Optional ) argument_description.
        * `cc_vtypes` - (Optional ) argument_description.
        * `hyp_props` - (Optional ) argument_description.
        * `info` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                    * `uuid` - argument_description.
    
