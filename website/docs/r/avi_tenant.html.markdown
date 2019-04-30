---
layout: "avi"
page_title: "Avi: avi_tenant"
sidebar_current: "docs-avi-resource-tenant"
description: |-
  Creates and manages Avi Tenant.
---

# avi_tenant

The Tenant resource allows the creation and management of Avi Tenant

## Example Usage

```hcl
resource "Tenant" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `config_settings` - (Optional ) argument_description.
        * `created_by` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `local` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                        * `uuid` - argument_description.
    
