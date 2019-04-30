---
layout: "avi"
page_title: "Avi: avi_applicationprofile"
sidebar_current: "docs-avi-resource-applicationprofile"
description: |-
  Creates and manages Avi ApplicationProfile.
---

# avi_applicationprofile

The ApplicationProfile resource allows the creation and management of Avi ApplicationProfile

## Example Usage

```hcl
resource "ApplicationProfile" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `cloud_config_cksum` - (Optional ) argument_description.
        * `created_by` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `dns_service_profile` - (Optional ) argument_description.
        * `dos_rl_profile` - (Optional ) argument_description.
        * `http_profile` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `preserve_client_ip` - (Optional ) argument_description.
        * `preserve_client_port` - (Optional ) argument_description.
        * `sip_service_profile` - (Optional ) argument_description.
        * `tcp_app_profile` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        * `type` - (Required) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                        * `uuid` - argument_description.
    
