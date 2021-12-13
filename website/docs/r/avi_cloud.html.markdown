<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
* `vtype` - (Required) Cloud type. Enum options - CLOUD_NONE, CLOUD_VCENTER, CLOUD_OPENSTACK, CLOUD_AWS, CLOUD_VCA, CLOUD_APIC, CLOUD_MESOS, CLOUD_LINUXSERVER, CLOUD_DOCKER_UCP, CLOUD_RANCHER, CLOUD_OSHIFT_K8S, CLOUD_AZURE, CLOUD_GCP, CLOUD_NSXT. Allowed in basic(allowed values- cloud_none,cloud_nsxt) edition, essentials(allowed values- cloud_none,cloud_vcenter) edition, enterprise edition.
* `autoscale_polling_interval` - (Optional) Cloudconnector polling interval in seconds for external autoscale groups, minimum 60 seconds. Allowed values are 60-3600. Field introduced in 18.2.2. Unit is seconds. Allowed in basic(allowed values- 60) edition, essentials(allowed values- 60) edition, enterprise edition.
* `aws_configuration` - (Optional) Dict settings for cloud.
* `azure_configuration` - (Optional) Field introduced in 17.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `cloudstack_configuration` - (Optional) Dict settings for cloud.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `custom_tags` - (Optional) Custom tags for all avi created resources in the cloud infrastructure. Field introduced in 17.1.5.
* `dhcp_enabled` - (Optional) Select the ip address management scheme.
* `dns_provider_ref` - (Optional) Dns profile for the cloud. It is a reference to an object of type ipamdnsproviderprofile.
* `dns_resolution_on_se` - (Optional) By default, pool member fqdns are resolved on the controller. When this is set, pool member fqdns are instead resolved on service engines in this cloud. This is useful in scenarios where pool member fqdns can only be resolved from service engines and not from the controller. Field introduced in 18.2.6. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `dns_resolvers` - (Optional) Dns resolver for the cloud. Field introduced in 20.1.5. Maximum of 1 items allowed.
* `docker_configuration` - (Optional) Dict settings for cloud.
* `east_west_dns_provider_ref` - (Optional) Dns profile for east-west services. It is a reference to an object of type ipamdnsproviderprofile.
* `east_west_ipam_provider_ref` - (Optional) Ipam profile for east-west services. Warning - please use virtual subnets in this ipam profile that do not conflict with the underlay networks or any overlay networks in the cluster. For example in aws and gcp, 169.254.0.0/16 is used for storing instance metadata. Hence, it should not be used in this profile. It is a reference to an object of type ipamdnsproviderprofile.
* `enable_vip_on_all_interfaces` - (Optional) Enable vip on all data interfaces for the cloud. Field introduced in 18.2.9, 20.1.1.
* `enable_vip_static_routes` - (Optional) Use static routes for vip side network resolution during virtualservice placement.
* `gcp_configuration` - (Optional) Google cloud platform configuration. Field introduced in 18.2.1. Allowed in basic edition, essentials edition, enterprise edition.
* `ip6_autocfg_enabled` - (Optional) Enable ipv6 auto configuration. Field introduced in 18.1.1.
* `ipam_provider_ref` - (Optional) Ipam profile for the cloud. It is a reference to an object of type ipamdnsproviderprofile.
* `license_tier` - (Optional) Specifies the default license tier which would be used by new se groups. This field by default inherits the value from system configuration. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS, ENTERPRISE_WITH_CLOUD_SERVICES. Field introduced in 17.2.5.
* `license_type` - (Optional) If no license type is specified then default license enforcement for the cloud type is chosen. The default mappings are container cloud is max ses, openstack and vmware is cores and linux it is sockets. Enum options - LIC_BACKEND_SERVERS, LIC_SOCKETS, LIC_CORES, LIC_HOSTS, LIC_SE_BANDWIDTH, LIC_METERED_SE_BANDWIDTH.
* `linuxserver_configuration` - (Optional) Dict settings for cloud.
* `maintenance_mode` - (Optional) Cloud is in maintenance mode. Field introduced in 20.1.7,21.1.3.
* `mtu` - (Optional) Mtu setting for the cloud. Unit is bytes.
* `nsxt_configuration` - (Optional) Nsx-t cloud platform configuration. Field introduced in 20.1.1. Allowed in essentials edition, enterprise edition.
* `obj_name_prefix` - (Optional) Default prefix for all automatically created objects in this cloud. This prefix can be overridden by the se-group template.
* `openstack_configuration` - (Optional) Dict settings for cloud.
* `prefer_static_routes` - (Optional) Prefer static routes over interface routes during virtualservice placement.
* `proxy_configuration` - (Optional) Dict settings for cloud.
* `rancher_configuration` - (Optional) Dict settings for cloud.
* `se_group_template_ref` - (Optional) The service engine group to use as template. It is a reference to an object of type serviceenginegroup. Field introduced in 18.2.5.
* `state_based_dns_registration` - (Optional) Dns records for vips are added/deleted based on the operational state of the vips. Field introduced in 17.1.12. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `vca_configuration` - (Optional) Dict settings for cloud.
* `vcenter_configuration` - (Optional) Dict settings for cloud.
* `vmc_deployment` - (Optional) This deployment is vmware on aws cloud. Field introduced in 20.1.5, 21.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

