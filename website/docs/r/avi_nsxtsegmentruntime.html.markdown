############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `cloud_ref` - (Optional) Nsxt segment belongs to cloud.
* `dhcp_enabled` - (Optional) Ip address management scheme for this segment associated network.
* `dhcp_ranges` - (Optional) Dhcp ranges configured in nsxt.
* `name` - (Optional) Segment object name.
* `nw_name` - (Optional) Network name.
* `nw_ref` - (Optional) Corresponding network object in avi.
* `opaque_network_id` - (Optional) Opaque network id.
* `segment_gw` - (Optional) Segment gateway.
* `segment_id` - (Optional) Segment id.
* `segname` - (Optional) Segment name.
* `subnet` - (Optional) Segment cidr.
* `tenant_ref` - (Optional) Nsxt segment belongs to tenant.
* `tier1_id` - (Optional) Tier1 router id.
* `vrf_context_ref` - (Optional) Corresponding vrf context object in avi.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid.

