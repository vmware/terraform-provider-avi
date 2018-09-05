---
layout: "avi"
page_title: "Avi: avi_ipaddrgroup"
sidebar_current: "docs-avi-resource-ipaddrgroup"
description: |-
Creates and manages Avi IpAddrGroup.
---

# avi_ipaddrgroup

The IpAddrGroup resource allows the creation and management of Avi IpAddrGroup

## Example Usage

```hcl
resource "IpAddrGroup" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `addrs` - (Optional ) argument_description.
        * `apic_epg_name` - (Optional ) argument_description.
        * `country_codes` - (Optional ) argument_description.
        * `description` - (Optional ) argument_description.
        * `ip_ports` - (Optional ) argument_description.
        * `marathon_app_name` - (Optional ) argument_description.
        * `marathon_service_port` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `prefixes` - (Optional ) argument_description.
        * `ranges` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
        
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                * `uuid` - argument_description.
    
