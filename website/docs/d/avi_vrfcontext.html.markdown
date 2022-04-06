<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_vrfcontext"
sidebar_current: "docs-avi-datasource-vrfcontext"
description: |-
  Get information of Avi VrfContext.
---

# avi_vrfcontext

This data source is used to to get avi_vrfcontext objects.

## Example Usage

```hcl
data "avi_vrfcontext" "foo_vrfcontext" {
    uuid = "vrfcontext-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search VrfContext by name.
* `uuid` - (Optional) Search VrfContext by uuid.
* `cloud_ref` - (Optional) Search VrfContext by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `attrs` - Key/value vrfcontext attributes. Field introduced in 20.1.2. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `bfd_profile` - Bfd configuration profile. Field introduced in 20.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `bgp_profile` - Bgp local and peer info. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `cloud_ref` - It is a reference to an object of type cloud. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `debugvrfcontext` - Configure debug flags for vrf. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `description` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `gateway_mon` - Configure ping based heartbeat check for gateway in service engines of vrf. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `internal_gateway_monitor` - Configure ping based heartbeat check for all default gateways in service engines of vrf. Field introduced in 17.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `lldp_enable` - Enable lldp. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise with any value edition, essentials(allowed values- true) edition, basic(allowed values- true) edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `static_routes` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `system_default` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

