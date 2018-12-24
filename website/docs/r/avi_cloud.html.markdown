---
layout: "avi"
page_title: "Avi: avi_cloud"
sidebar_current: "docs-avi-resource-cloud"
description: |-
  Creates and manages Avi Cloud.
---

# avi_cloud

The Cloud resource allows the creation and management of Avi Cloud

## Example Usage

```hcl
resource "Cloud" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `apic_configuration` - (Optional ) argument_description.
        * `apic_mode` - (Optional ) argument_description.
        * `aws_configuration` - (Optional ) argument_description.
        * `azure_configuration` - (Optional ) argument_description.
        * `cloudstack_configuration` - (Optional ) argument_description.
        * `custom_tags` - (Optional ) argument_description.
        * `dhcp_enabled` - (Optional ) argument_description.
        * `dns_provider_ref` - (Optional ) argument_description.
        * `docker_configuration` - (Optional ) argument_description.
        * `east_west_dns_provider_ref` - (Optional ) argument_description.
        * `east_west_ipam_provider_ref` - (Optional ) argument_description.
        * `enable_vip_static_routes` - (Optional ) argument_description.
        * `gcp_configuration` - (Optional ) argument_description.
        * `ip6_autocfg_enabled` - (Optional ) argument_description.
        * `ipam_provider_ref` - (Optional ) argument_description.
        * `license_tier` - (Optional ) argument_description.
        * `license_type` - (Optional ) argument_description.
        * `linuxserver_configuration` - (Optional ) argument_description.
        * `mesos_configuration` - (Optional ) argument_description.
        * `mtu` - (Optional ) argument_description.
        * `name` - (Required) argument_description.
        * `nsx_configuration` - (Optional ) argument_description.
        * `obj_name_prefix` - (Optional ) argument_description.
        * `openstack_configuration` - (Optional ) argument_description.
        * `oshiftk8s_configuration` - (Optional ) argument_description.
        * `prefer_static_routes` - (Optional ) argument_description.
        * `proxy_configuration` - (Optional ) argument_description.
        * `rancher_configuration` - (Optional ) argument_description.
        * `state_based_dns_registration` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
            * `vca_configuration` - (Optional ) argument_description.
        * `vcenter_configuration` - (Optional ) argument_description.
        * `vtype` - (Required) argument_description.
    
### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                                                                                                                            * `uuid` - argument_description.
                
