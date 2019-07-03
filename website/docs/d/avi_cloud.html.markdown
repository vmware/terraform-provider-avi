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
data "avi_cloud" "foo_cloud" {
    uuid = "cloud-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Cloud by name.
* `uuid` - (Optional) Search Cloud by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `apic_configuration` - Dict settings for cloud.
* `apic_mode` - Boolean flag to set apic_mode.
* `autoscale_polling_interval` - Cloudconnector polling interval in seconds for external autoscale groups, minimum 60 seconds.
* `aws_configuration` - Dict settings for cloud.
* `azure_configuration` - Field introduced in 17.2.1.
* `cloudstack_configuration` - Dict settings for cloud.
* `custom_tags` - Custom tags for all avi created resources in the cloud infrastructure.
* `dhcp_enabled` - Select the ip address management scheme.
* `dns_provider_ref` - Dns profile for the cloud.
* `docker_configuration` - Dict settings for cloud.
* `east_west_dns_provider_ref` - Dns profile for east-west services.
* `east_west_ipam_provider_ref` - Ipam profile for east-west services.
* `enable_vip_static_routes` - Use static routes for vip side network resolution during virtualservice placement.
* `gcp_configuration` - Google cloud platform configuration.
* `ip6_autocfg_enabled` - Enable ipv6 auto configuration.
* `ipam_provider_ref` - Ipam profile for the cloud.
* `license_tier` - Specifies the default license tier which would be used by new se groups.
* `license_type` - If no license type is specified then default license enforcement for the cloud type is chosen.
* `linuxserver_configuration` - Dict settings for cloud.
* `mtu` - Mtu setting for the cloud.
* `name` - Name of the object.
* `nsx_configuration` - Configuration parameters for nsx manager.
* `obj_name_prefix` - Default prefix for all automatically created objects in this cloud.
* `openstack_configuration` - Dict settings for cloud.
* `oshiftk8s_configuration` - Dict settings for cloud.
* `prefer_static_routes` - Prefer static routes over interface routes during virtualservice placement.
* `proxy_configuration` - Dict settings for cloud.
* `rancher_configuration` - Dict settings for cloud.
* `se_group_template_ref` - The service engine group to use as template.
* `state_based_dns_registration` - Dns records for vips are added/deleted based on the operational state of the vips.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.
* `vca_configuration` - Dict settings for cloud.
* `vcenter_configuration` - Dict settings for cloud.
* `vtype` - Cloud type.

