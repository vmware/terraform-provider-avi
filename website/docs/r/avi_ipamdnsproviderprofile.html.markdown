---
layout: "avi"
page_title: "Avi: avi_ipamdnsproviderprofile"
sidebar_current: "docs-avi-resource-ipamdnsproviderprofile"
description: |-
Creates and manages Avi IpamDnsProviderProfile.
---

# avi_ipamdnsproviderprofile

The IpamDnsProviderProfile resource allows the creation and management of Avi IpamDnsProviderProfile

## Example Usage

```hcl
resource "IpamDnsProviderProfile" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `allocate_ip_in_vrf` - (Optional ) argument_description.
        * `aws_profile` - (Optional ) argument_description.
        * `azure_profile` - (Optional ) argument_description.
        * `custom_profile` - (Optional ) argument_description.
        * `gcp_profile` - (Optional ) argument_description.
        * `infoblox_profile` - (Optional ) argument_description.
        * `internal_profile` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `oci_profile` - (Optional ) argument_description.
        * `openstack_profile` - (Optional ) argument_description.
        * `proxy_configuration` - (Optional ) argument_description.
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
    
