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
resource "avi_cloud" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `vtype` - (Required) Cloud type.
* `apic_configuration` - (Optional) Dict settings for cloud.
* `apic_mode` - (Optional) Boolean flag to set apic_mode.
* `autoscale_polling_interval` - (Optional) Cloudconnector polling interval in seconds for external autoscale groups, minimum 60 seconds.
* `aws_configuration` - (Optional) Dict settings for cloud.
* `azure_configuration` - (Optional) Field introduced in 17.2.1.
* `cloudstack_configuration` - (Optional) Dict settings for cloud.
* `custom_tags` - (Optional) Custom tags for all avi created resources in the cloud infrastructure.
* `dhcp_enabled` - (Optional) Select the ip address management scheme.
* `dns_provider_ref` - (Optional) Dns profile for the cloud.
* `docker_configuration` - (Optional) Dict settings for cloud.
* `east_west_dns_provider_ref` - (Optional) Dns profile for east-west services.
* `east_west_ipam_provider_ref` - (Optional) Ipam profile for east-west services.
* `enable_vip_static_routes` - (Optional) Use static routes for vip side network resolution during virtualservice placement.
* `gcp_configuration` - (Optional) Google cloud platform configuration.
* `ip6_autocfg_enabled` - (Optional) Enable ipv6 auto configuration.
* `ipam_provider_ref` - (Optional) Ipam profile for the cloud.
* `license_tier` - (Optional) Specifies the default license tier which would be used by new se groups.
* `license_type` - (Optional) If no license type is specified then default license enforcement for the cloud type is chosen.
* `linuxserver_configuration` - (Optional) Dict settings for cloud.
* `mtu` - (Optional) Mtu setting for the cloud.
* `nsx_configuration` - (Optional) Configuration parameters for nsx manager.
* `obj_name_prefix` - (Optional) Default prefix for all automatically created objects in this cloud.
* `openstack_configuration` - (Optional) Dict settings for cloud.
* `oshiftk8s_configuration` - (Optional) Dict settings for cloud.
* `prefer_static_routes` - (Optional) Prefer static routes over interface routes during virtualservice placement.
* `proxy_configuration` - (Optional) Dict settings for cloud.
* `rancher_configuration` - (Optional) Dict settings for cloud.
* `se_group_template_ref` - (Optional) The service engine group to use as template.
* `state_based_dns_registration` - (Optional) Dns records for vips are added/deleted based on the operational state of the vips.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `vca_configuration` - (Optional) Dict settings for cloud.
* `vcenter_configuration` - (Optional) Dict settings for cloud.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

