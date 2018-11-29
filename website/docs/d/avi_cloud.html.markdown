---
layout: "avi"
page_title: "AVI: avi_cloud"
sidebar_current: "docs-avi-datasource-cloud"
description: |-
  Get information of Avi Cloud.
---

# avi_cloud

This data source is used to to get avi_cloud objects.

## Example Usage

```hcl
data "Cloud" "foo_Cloud" {
    uuid = "Cloud-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Cloud by name.
* `uuid` - (Optional) Search Cloud by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `apic_configuration` - General description.
* `apic_mode` - General description.
* `aws_configuration` - General description.
* `azure_configuration` - Field introduced in 17.2.1.
* `cloudstack_configuration` - General description.
* `custom_tags` - Custom tags for all avi created resources in the cloud infrastructure.
* `dhcp_enabled` - Select the ip address management scheme.
* `dns_provider_ref` - Dns profile for the cloud.
* `docker_configuration` - General description.
* `east_west_dns_provider_ref` - Dns profile for east-west services.
* `east_west_ipam_provider_ref` - Ipam profile for east-west services.
* `enable_vip_static_routes` - Use static routes for vip side network resolution during virtualservice placement.
* `ip6_autocfg_enabled` - Enable ipv6 auto configuration.
* `ipam_provider_ref` - Ipam profile for the cloud.
* `license_tier` - Specifies the default license tier which would be used by new se groups.
* `license_type` - If no license type is specified then default license enforcement for the cloud type is chosen.
* `linuxserver_configuration` - General description.
* `mesos_configuration` - General description.
* `mtu` - Mtu setting for the cloud.
* `name` - General description.
* `nsx_configuration` - Configuration parameters for nsx manager.
* `obj_name_prefix` - Default prefix for all automatically created objects in this cloud.
* `openstack_configuration` - General description.
* `oshiftk8s_configuration` - General description.
* `prefer_static_routes` - Prefer static routes over interface routes during virtualservice placement.
* `proxy_configuration` - General description.
* `rancher_configuration` - General description.
* `state_based_dns_registration` - Dns records for vips are added/deleted based on the operational state of the vips.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.
* `vca_configuration` - General description.
* `vcenter_configuration` - General description.
* `vtype` - Cloud type.
