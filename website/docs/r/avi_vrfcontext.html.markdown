<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_vrfcontext"
sidebar_current: "docs-avi-resource-vrfcontext"
description: |-
  Creates and manages Avi VrfContext.
---

# avi_vrfcontext

The VrfContext resource allows the creation and management of Avi VrfContext

## Example Usage

```hcl
resource "avi_vrfcontext" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `attrs` - (Optional) Key/value vrfcontext attributes. Field introduced in 20.1.2. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `bfd_profile` - (Optional) Bfd configuration profile. Field introduced in 20.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `bgp_profile` - (Optional) Bgp local and peer info. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `debugvrfcontext` - (Optional) Configure debug flags for vrf. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `gateway_mon` - (Optional) Configure ping based heartbeat check for gateway in service engines of vrf. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `internal_gateway_monitor` - (Optional) Configure ping based heartbeat check for all default gateways in service engines of vrf. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `lldp_enable` - (Optional) Enable lldp. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise edition with any value, essentials edition(allowed values- true), basic edition(allowed values- true), enterprise with cloud services edition.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `static_routes` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `system_default` - (Optional) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

