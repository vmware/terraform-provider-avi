---
layout: "avi"
page_title: "Avi: avi_gslbgeodbprofile"
sidebar_current: "docs-avi-resource-gslbgeodbprofile"
description: |-
Creates and manages Avi GslbGeoDbProfile.
---

# avi_gslbgeodbprofile

The GslbGeoDbProfile resource allows the creation and management of Avi GslbGeoDbProfile

## Example Usage

```hcl
resource "GslbGeoDbProfile" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `description` - (Optional ) argument_description.
        * `entries` - (Optional ) argument_description.
        * `is_federated` - (Optional ) argument_description.
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
    
