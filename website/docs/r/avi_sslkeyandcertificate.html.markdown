---
layout: "avi"
page_title: "Avi: avi_sslkeyandcertificate"
sidebar_current: "docs-avi-resource-sslkeyandcertificate"
description: |-
  Creates and manages Avi SSLKeyAndCertificate.
---

# avi_sslkeyandcertificate

The SSLKeyAndCertificate resource allows the creation and management of Avi SSLKeyAndCertificate

## Example Usage

```hcl
resource "SSLKeyAndCertificate" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `ca_certs` - (Optional ) argument_description.
        * `certificate` - (Required) argument_description.
        * `certificate_base64` - (Optional ) argument_description.
        * `certificate_management_profile_ref` - (Optional ) argument_description.
        * `created_by` - (Optional ) argument_description.
        * `dynamic_params` - (Optional ) argument_description.
        * `enckey_base64` - (Optional ) argument_description.
        * `enckey_name` - (Optional ) argument_description.
        * `format` - (Optional ) argument_description.
        * `hardwaresecuritymodulegroup_ref` - (Optional ) argument_description.
        * `key` - (Optional ) argument_description.
        * `key_base64` - (Optional ) argument_description.
        * `key_params` - (Optional ) argument_description.
        * `key_passphrase` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `status` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `type` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                            * `uuid` - argument_description.
    
