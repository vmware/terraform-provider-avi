---
layout: "avi"
page_title: "Avi: avi_vsvip"
sidebar_current: "docs-avi-resource-vsvip"
description: |-
  Creates and manages Avi VsVip.
---

# avi_vsvip

The VsVip resource allows the creation and management of Avi VsVip

## Example Usage

```hcl
resource "avi_vsvip" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the vsvip object.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `dns_info` - (Optional) Service discovery specific data including fully qualified domain name, type and time-to-live of the dns record.
* `east_west_placement` - (Optional) Force placement on all service engines in the service engine group (container clouds only).
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `use_standard_alb` - (Optional) This overrides the cloud level default and needs to match the se group value in which it will be used if the se group use_standard_alb value is set.
* `vip` - (Optional) List of virtual service ips and other shareable entities.
* `vrf_context_ref` - (Optional) Virtual routing context that the virtual service is bound to.
* `vsvip_cloud_config_cksum` - (Optional) Checksum of cloud configuration for vsvip.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the vsvip object.

