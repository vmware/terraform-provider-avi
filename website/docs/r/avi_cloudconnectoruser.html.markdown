---
layout: "avi"
page_title: "Avi: avi_cloudconnectoruser"
sidebar_current: "docs-avi-resource-cloudconnectoruser"
description: |-
  Creates and manages Avi CloudConnectorUser.
---

# avi_cloudconnectoruser

The CloudConnectorUser resource allows the creation and management of Avi CloudConnectorUser

## Example Usage

```hcl
resource "CloudConnectorUser" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `azure_serviceprincipal` - (Optional ) argument_description.
        * `azure_userpass` - (Optional ) argument_description.
        * `gcp_credentials` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `oci_credentials` - (Optional ) argument_description.
        * `password` - (Optional ) argument_description.
        * `private_key` - (Optional ) argument_description.
        * `public_key` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `tencent_credentials` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                            * `uuid` - argument_description.
    
