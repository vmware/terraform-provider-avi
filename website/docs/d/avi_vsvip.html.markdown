<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_vsvip"
sidebar_current: "docs-avi-datasource-vsvip"
description: |-
  Get information of Avi VsVip.
---

# avi_vsvip

This data source is used to to get avi_vsvip objects.

## Example Usage

```hcl
data "avi_vsvip" "foo_vsvip" {
    uuid = "vsvip-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search VsVip by name.
* `uuid` - (Optional) Search VsVip by uuid.
* `cloud_ref` - (Optional) Search VsVip by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `bgp_local_preference` - Local_pref to be used for vsvip advertised. Applicable only over ibgp. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `bgp_num_as_path_prepend` - Number of times the local as should be prepended additionally to vsvip. Applicable only over ebgp. Allowed values are 1-10. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `bgp_peer_labels` - Select bgp peers, using peer label, for vsvip advertisement. Field introduced in 20.1.5. Maximum of 128 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `cloud_ref` - It is a reference to an object of type cloud. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `dns_info` - Service discovery specific data including fully qualified domain name, type and time-to-live of the dns record. Field introduced in 17.1.1. Maximum of 1000 items allowed. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `east_west_placement` - Force placement on all service engines in the service engine group (container clouds only). Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic edition(allowed values- false), enterprise with cloud services edition.
* `ipam_selector` - Determines the set of ipam networks to use for this vsvip. Selector type must be selector_ipam and only one label is supported. Field introduced in 20.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - Name for the vsvip object. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tier1_lr` - This sets the placement scope of virtualservice to given tier1 logical router in nsx-t. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `use_standard_alb` - This overrides the cloud level default and needs to match the se group value in which it will be used if the se group use_standard_alb value is set. This is only used when fip is used for vs on azure cloud. Field introduced in 18.2.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Uuid of the vsvip object. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vip` - List of virtual service ips and other shareable entities. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vrf_context_ref` - Virtual routing context that the virtual service is bound to. This is used to provide the isolation of the set of networks the application is attached to. It is a reference to an object of type vrfcontext. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vsvip_cloud_config_cksum` - Checksum of cloud configuration for vsvip. Internally set by cloud connector. Field introduced in 17.2.9, 18.1.2. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

