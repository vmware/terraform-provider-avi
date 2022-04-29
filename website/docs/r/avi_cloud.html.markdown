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

* `name` - (Required) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vtype` - (Required) Cloud type. Enum options - CLOUD_NONE, CLOUD_VCENTER, CLOUD_OPENSTACK, CLOUD_AWS, CLOUD_VCA, CLOUD_APIC, CLOUD_MESOS, CLOUD_LINUXSERVER, CLOUD_DOCKER_UCP, CLOUD_RANCHER, CLOUD_OSHIFT_K8S, CLOUD_AZURE, CLOUD_GCP, CLOUD_NSXT. Allowed in enterprise edition with any value, essentials edition(allowed values- cloud_none,cloud_vcenter), basic edition(allowed values- cloud_none,cloud_nsxt), enterprise with cloud services edition.
* `autoscale_polling_interval` - (Optional) Cloudconnector polling interval in seconds for external autoscale groups, minimum 60 seconds. Allowed values are 60-3600. Field introduced in 18.2.2. Unit is seconds. Allowed in enterprise edition with any value, essentials edition(allowed values- 60), basic edition(allowed values- 60), enterprise with cloud services edition.
* `aws_configuration` - (Optional) Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `azure_configuration` - (Optional) Field introduced in 17.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `cloudstack_configuration` - (Optional) Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `custom_tags` - (Optional) Custom tags for all avi created resources in the cloud infrastructure. Field introduced in 17.1.5. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dhcp_enabled` - (Optional) Select the ip address management scheme. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dns_provider_ref` - (Optional) Dns profile for the cloud. It is a reference to an object of type ipamdnsproviderprofile. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `dns_resolution_on_se` - (Optional) By default, pool member fqdns are resolved on the controller. When this is set, pool member fqdns are instead resolved on service engines in this cloud. This is useful in scenarios where pool member fqdns can only be resolved from service engines and not from the controller. Field introduced in 18.2.6. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `dns_resolvers` - (Optional) Dns resolver for the cloud. Field introduced in 20.1.5. Maximum of 1 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `docker_configuration` - (Optional) Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `east_west_dns_provider_ref` - (Optional) Dns profile for east-west services. It is a reference to an object of type ipamdnsproviderprofile. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `east_west_ipam_provider_ref` - (Optional) Ipam profile for east-west services. Warning - please use virtual subnets in this ipam profile that do not conflict with the underlay networks or any overlay networks in the cluster. For example in aws and gcp, 169.254.0.0/16 is used for storing instance metadata. Hence, it should not be used in this profile. It is a reference to an object of type ipamdnsproviderprofile. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `enable_vip_on_all_interfaces` - (Optional) Enable vip on all data interfaces for the cloud. Field introduced in 18.2.9, 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_vip_static_routes` - (Optional) Use static routes for vip side network resolution during virtualservice placement. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `gcp_configuration` - (Optional) Google cloud platform configuration. Field introduced in 18.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `ip6_autocfg_enabled` - (Optional) Enable ipv6 auto configuration. Field introduced in 18.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `ipam_provider_ref` - (Optional) Ipam profile for the cloud. It is a reference to an object of type ipamdnsproviderprofile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `license_tier` - (Optional) Specifies the default license tier which would be used by new se groups. This field by default inherits the value from system configuration. Enum options - ENTERPRISE_16, ENTERPRISE, ENTERPRISE_18, BASIC, ESSENTIALS, ENTERPRISE_WITH_CLOUD_SERVICES. Field introduced in 17.2.5. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `license_type` - (Optional) If no license type is specified then default license enforcement for the cloud type is chosen. The default mappings are container cloud is max ses, openstack and vmware is cores and linux it is sockets. Enum options - LIC_BACKEND_SERVERS, LIC_SOCKETS, LIC_CORES, LIC_HOSTS, LIC_SE_BANDWIDTH, LIC_METERED_SE_BANDWIDTH. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `linuxserver_configuration` - (Optional) Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `maintenance_mode` - (Optional) Cloud is in maintenance mode. Field introduced in 20.1.7,21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `mtu` - (Optional) Mtu setting for the cloud. Unit is bytes. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `nsxt_configuration` - (Optional) Nsx-t cloud platform configuration. Field introduced in 20.1.1. Allowed in enterprise edition with any value, basic, enterprise with cloud services edition.
* `obj_name_prefix` - (Optional) Default prefix for all automatically created objects in this cloud. This prefix can be overridden by the se-group template. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `openstack_configuration` - (Optional) Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `prefer_static_routes` - (Optional) Prefer static routes over interface routes during virtualservice placement. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `proxy_configuration` - (Optional) Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `rancher_configuration` - (Optional) Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `se_group_template_ref` - (Optional) The service engine group to use as template. It is a reference to an object of type serviceenginegroup. Field introduced in 18.2.5. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `state_based_dns_registration` - (Optional) Dns records for vips are added/deleted based on the operational state of the vips. Field introduced in 17.1.12. Allowed in enterprise edition with any value, essentials edition(allowed values- true), basic edition(allowed values- true), enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vca_configuration` - (Optional) Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vcenter_configuration` - (Optional) Allowed in enterprise edition with any value, essentials, enterprise with cloud services edition.
* `vmc_deployment` - (Optional) This deployment is vmware on aws cloud. Field introduced in 20.1.5, 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

