<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
* `apic_mode` - Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `autoscale_polling_interval` - Cloudconnector polling interval in seconds for external autoscale groups, minimum 60 seconds. Allowed values are 60-3600. Field introduced in 18.2.2. Unit is seconds. Allowed in basic(allowed values- 60) edition, essentials(allowed values- 60) edition, enterprise edition.
* `aws_configuration` - Dict settings for cloud.
* `azure_configuration` - Field introduced in 17.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `cloudstack_configuration` - Dict settings for cloud.
* `custom_tags` - Custom tags for all avi created resources in the cloud infrastructure. Field introduced in 17.1.5.
* `dhcp_enabled` - Select the ip address management scheme.
* `dns_provider_ref` - Dns profile for the cloud. It is a reference to an object of type ipamdnsproviderprofile.
* `dns_resolution_on_se` - By default, pool member fqdns are resolved on the controller. When this is set, pool member fqdns are instead resolved on service engines in this cloud. This is useful in scenarios where pool member fqdns can only be resolved from service engines and not from the controller. Field introduced in 18.2.6. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `dns_resolvers` - Dns resolver for the cloud. Field introduced in 20.1.5. Maximum of 1 items allowed.
* `docker_configuration` - Dict settings for cloud.
* `east_west_dns_provider_ref` - Dns profile for east-west services. It is a reference to an object of type ipamdnsproviderprofile.
* `east_west_ipam_provider_ref` - Ipam profile for east-west services. Warning - please use virtual subnets in this ipam profile that do not conflict with the underlay networks or any overlay networks in the cluster. For example in aws and gcp, 169.254.0.0/16 is used for storing instance metadata. Hence, it should not be used in this profile. It is a reference to an object of type ipamdnsproviderprofile.
* `enable_vip_on_all_interfaces` - Enable vip on all data interfaces for the cloud. Field introduced in 18.2.9, 20.1.1.
* `enable_vip_static_routes` - Use static routes for vip side network resolution during virtualservice placement.
* `gcp_configuration` - Google cloud platform configuration. Field introduced in 18.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `ip6_autocfg_enabled` - Enable ipv6 auto configuration. Field introduced in 18.1.1.
* `ipam_provider_ref` - Ipam profile for the cloud. It is a reference to an object of type ipamdnsproviderprofile.
* `license_tier` - Specifies the default license tier which would be used by new se groups. This field by default inherits the value from system configuration. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS. Field introduced in 17.2.5.
* `license_type` - If no license type is specified then default license enforcement for the cloud type is chosen. The default mappings are container cloud is max ses, openstack and vmware is cores and linux it is sockets. Enum options - LIC_BACKEND_SERVERS, LIC_SOCKETS, LIC_CORES, LIC_HOSTS, LIC_SE_BANDWIDTH, LIC_METERED_SE_BANDWIDTH.
* `linuxserver_configuration` - Dict settings for cloud.
* `mtu` - Mtu setting for the cloud. Unit is bytes.
* `name` - Name of the object.
* `nsx_configuration` - Configuration parameters for nsx manager. Field introduced in 17.1.1.
* `nsxt_configuration` - Nsx-t cloud platform configuration. Field introduced in 20.1.1. Allowed in essentials edition, enterprise edition.
* `obj_name_prefix` - Default prefix for all automatically created objects in this cloud. This prefix can be overridden by the se-group template.
* `openstack_configuration` - Dict settings for cloud.
* `prefer_static_routes` - Prefer static routes over interface routes during virtualservice placement.
* `proxy_configuration` - Dict settings for cloud.
* `rancher_configuration` - Dict settings for cloud.
* `se_group_template_ref` - The service engine group to use as template. It is a reference to an object of type serviceenginegroup. Field introduced in 18.2.5.
* `state_based_dns_registration` - Dns records for vips are added/deleted based on the operational state of the vips. Field introduced in 17.1.12. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.
* `vca_configuration` - Dict settings for cloud.
* `vcenter_configuration` - Dict settings for cloud.
* `vmc_deployment` - This deployment is vmware on aws cloud. Field introduced in 20.1.5.
* `vtype` - Cloud type. Enum options - CLOUD_NONE, CLOUD_VCENTER, CLOUD_OPENSTACK, CLOUD_AWS, CLOUD_VCA, CLOUD_APIC, CLOUD_MESOS, CLOUD_LINUXSERVER, CLOUD_DOCKER_UCP, CLOUD_RANCHER, CLOUD_OSHIFT_K8S, CLOUD_AZURE, CLOUD_GCP, CLOUD_NSXT. Allowed in basic(allowed values- cloud_none,cloud_nsxt) edition, essentials(allowed values- cloud_none,cloud_vcenter) edition, enterprise edition.

