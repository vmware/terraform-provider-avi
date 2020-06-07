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

* `cloud_ref` - Nsxt segment belongs to cloud.
* `dhcp_enabled` - Ip address management scheme for this segment associated network.
* `dhcp_ranges` - Dhcp ranges configured in nsxt.
* `name` - Segment object name.
* `nw_name` - Network name.
* `nw_ref` - Corresponding network object in avi.
* `opaque_network_id` - Opaque network id.
* `segment_gw` - Segment gateway.
* `segment_id` - Segment id.
* `segname` - Segment name.
* `subnet` - Segment cidr.
* `tenant_ref` - Nsxt segment belongs to tenant.
* `tier1_id` - Tier1 router id.
* `uuid` - Uuid.
* `vrf_context_ref` - Corresponding vrf context object in avi.

