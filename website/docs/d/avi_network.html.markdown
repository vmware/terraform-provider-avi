<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_network"
sidebar_current: "docs-avi-datasource-network"
description: |-
  Get information of Avi Network.
---

# avi_network

This data source is used to to get avi_network objects.

## Example Usage

```hcl
data "avi_network" "foo_network" {
    uuid = "network-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search Network by name.
* `uuid` - (Optional) Search Network by uuid.
* `cloud_ref` - (Optional) Search Network by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `attrs` - Key/value network attributes. Field introduced in 20.1.1.
* `cloud_ref` - It is a reference to an object of type cloud.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `configured_subnets` - List of list.
* `dhcp_enabled` - Select the ip address management scheme for this network.
* `exclude_discovered_subnets` - When selected, excludes all discovered subnets in this network from consideration for virtual service placement.
* `ip6_autocfg_enabled` - Enable ipv6 auto configuration. Field introduced in 18.1.1.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name of the object.
* `synced_from_se` - Boolean flag to set synced_from_se.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.
* `vcenter_dvs` - Boolean flag to set vcenter_dvs.
* `vrf_context_ref` - It is a reference to an object of type vrfcontext.

