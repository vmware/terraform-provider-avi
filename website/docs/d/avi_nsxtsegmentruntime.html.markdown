<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_nsxtsegmentruntime"
sidebar_current: "docs-avi-datasource-nsxtsegmentruntime"
description: |-
  Get information of Avi NsxtSegmentRuntime.
---

# avi_nsxtsegmentruntime

This data source is used to to get avi_nsxtsegmentruntime objects.

## Example Usage

```hcl
data "avi_nsxtsegmentruntime" "foo_nsxtsegmentruntime" {
    uuid = "nsxtsegmentruntime-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search NsxtSegmentRuntime by name.
* `uuid` - (Optional) Search NsxtSegmentRuntime by uuid.
* `cloud_ref` - (Optional) Search NsxtSegmentRuntime by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_ref` - Nsxt segment belongs to cloud. It is a reference to an object of type cloud. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `dhcp6_ranges` - V6 dhcp ranges configured in nsxt. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `dhcp_enabled` - Ip address management scheme for this segment associated network. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `dhcp_ranges` - Dhcp ranges configured in nsxt. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `name` - Segment object name. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `nw_name` - Network name. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `nw_ref` - Corresponding network object in avi. It is a reference to an object of type network. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `opaque_network_id` - Opaque network id. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `segment_gw` - Segment gateway. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `segment_gw6` - V6 segment gateway. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `segment_id` - Segment id. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `segname` - Segment name. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `subnet` - Segment cidr. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `subnet6` - V6 segment cidr. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - Nsxt segment belongs to tenant. It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tier1_id` - Tier1 router id. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Uuid. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `vlan_ids` - Segment vlan ids. Field introduced in 20.1.5. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `vrf_context_ref` - Corresponding vrf context object in avi. It is a reference to an object of type vrfcontext. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

