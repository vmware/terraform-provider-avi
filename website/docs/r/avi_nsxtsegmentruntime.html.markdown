<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_nsxtsegmentruntime"
sidebar_current: "docs-avi-resource-nsxtsegmentruntime"
description: |-
  Creates and manages Avi NsxtSegmentRuntime.
---

# avi_nsxtsegmentruntime

The NsxtSegmentRuntime resource allows the creation and management of Avi NsxtSegmentRuntime

## Example Usage

```hcl
resource "avi_nsxtsegmentruntime" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `cloud_ref` - (Optional) Nsxt segment belongs to cloud. It is a reference to an object of type cloud. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dhcp6_ranges` - (Optional) V6 dhcp ranges configured in nsxt. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dhcp_enabled` - (Optional) Ip address management scheme for this segment associated network. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dhcp_ranges` - (Optional) Dhcp ranges configured in nsxt. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - (Optional) Segment object name. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `nw_name` - (Optional) Network name. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `nw_ref` - (Optional) Corresponding network object in avi. It is a reference to an object of type network. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `opaque_network_id` - (Optional) Opaque network id. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `segment_gw` - (Optional) Segment gateway. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `segment_gw6` - (Optional) V6 segment gateway. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `segment_id` - (Optional) Segment id. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `segname` - (Optional) Segment name. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `subnet` - (Optional) Segment cidr. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `subnet6` - (Optional) V6 segment cidr. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Nsxt segment belongs to tenant. It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tier1_id` - (Optional) Tier1 router id. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `vlan_ids` - (Optional) Segment vlan ids. Field introduced in 20.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vrf_context_ref` - (Optional) Corresponding vrf context object in avi. It is a reference to an object of type vrfcontext. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

