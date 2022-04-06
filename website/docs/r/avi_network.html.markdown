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

* `name` - (Required) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `attrs` - (Optional) Key/value network attributes. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `configured_subnets` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `dhcp_enabled` - (Optional) Select the ip address management scheme for this network. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `exclude_discovered_subnets` - (Optional) When selected, excludes all discovered subnets in this network from consideration for virtual service placement. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `ip6_autocfg_enabled` - (Optional) Enable ipv6 auto configuration. Field introduced in 18.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `synced_from_se` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `vcenter_dvs` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `vrf_context_ref` - (Optional) It is a reference to an object of type vrfcontext. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

