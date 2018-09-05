---
layout: "avi"
page_title: "Avi: avi_networksecuritypolicy"
sidebar_current: "docs-avi-resource-networksecuritypolicy"
description: |-
Creates and manages Avi NetworkSecurityPolicy.
---

# avi_networksecuritypolicy

The NetworkSecurityPolicy resource allows the creation and management of Avi NetworkSecurityPolicy

## Example Usage

```hcl
resource "NetworkSecurityPolicy" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `cloud_config_cksum` - (Optional ) argument_description.
        * `created_by` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `name` - (Optional ) argument_description.
        * `rules` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                            * `uuid` - argument_description.
    
