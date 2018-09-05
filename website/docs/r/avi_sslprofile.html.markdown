---
layout: "avi"
page_title: "Avi: avi_sslprofile"
sidebar_current: "docs-avi-resource-sslprofile"
description: |-
Creates and manages Avi SSLProfile.
---

# avi_sslprofile

The SSLProfile resource allows the creation and management of Avi SSLProfile

## Example Usage

```hcl
resource "SSLProfile" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `accepted_ciphers` - (Optional ) argument_description.
        * `accepted_versions` - (Optional ) argument_description.
        * `cipher_enums` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `dhparam` - (Optional ) argument_description.
        * `enable_ssl_session_reuse` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `prefer_client_cipher_ordering` - (Optional ) argument_description.
        * `send_close_notify` - (Optional ) argument_description.
        * `ssl_rating` - (Optional ) argument_description.
        * `ssl_session_timeout` - (Optional ) argument_description.
        * `tags` - (Optional ) argument_description.
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
    
