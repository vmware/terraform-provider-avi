<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_network"
sidebar_current: "docs-avi-resource-network"
description: |-
  Creates and manages Avi Network.
---

# avi_network

The Network resource allows the creation and management of Avi Network

## Example Usage

```hcl
resource "avi_network" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `attrs` - (Optional) Key/value network attributes. Field introduced in 20.1.1.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `configured_subnets` - (Optional) List of list.
* `dhcp_enabled` - (Optional) Select the ip address management scheme for this network.
* `exclude_discovered_subnets` - (Optional) When selected, excludes all discovered subnets in this network from consideration for virtual service placement.
* `ip6_autocfg_enabled` - (Optional) Enable ipv6 auto configuration. Field introduced in 18.1.1.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `synced_from_se` - (Optional) Boolean flag to set synced_from_se.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `vcenter_dvs` - (Optional) Boolean flag to set vcenter_dvs.
* `vrf_context_ref` - (Optional) It is a reference to an object of type vrfcontext.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

