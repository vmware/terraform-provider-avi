---
layout: "avi"
page_title: "Avi: avi_network"
sidebar_current: "docs-avi-resource-network"
description: |-
Creates and manages Avi Network.
---

# avi_network

The Network resource allows the creation and management of Avi Network

## Example Usage

```hcl
resource "Network" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `cloud_ref` - (Optional ) argument_description.
        * `configured_subnets` - (Optional ) argument_description.
        * `dhcp_enabled` - (Optional ) argument_description.
        * `exclude_discovered_subnets` - (Optional ) argument_description.
        * `ip6_autocfg_enabled` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `synced_from_se` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
            * `vcenter_dvs` - (Optional ) argument_description.
        * `vrf_context_ref` - (Optional ) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                    * `uuid` - argument_description.
            
