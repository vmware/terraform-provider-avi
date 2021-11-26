############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

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

* `attrs` - Key/value vrfcontext attributes. Field introduced in 20.1.2. Allowed in basic edition, essentials edition, enterprise edition.
* `bfd_profile` - Bfd configuration profile. Field introduced in 20.1.1. Allowed in basic edition, essentials edition, enterprise edition.
* `bgp_profile` - Bgp local and peer info.
* `cloud_ref` - It is a reference to an object of type cloud.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `debugvrfcontext` - Configure debug flags for vrf. Field introduced in 17.1.1.
* `description` - User defined description for the object.
* `gateway_mon` - Configure ping based heartbeat check for gateway in service engines of vrf.
* `internal_gateway_monitor` - Configure ping based heartbeat check for all default gateways in service engines of vrf. Field introduced in 17.1.1.
* `lldp_enable` - Enable lldp. Field introduced in 18.2.10, 20.1.1. Allowed in basic(allowed values- true) edition, essentials(allowed values- true) edition, enterprise edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name of the object.
* `static_routes` - List of list.
* `system_default` - Boolean flag to set system_default.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.

