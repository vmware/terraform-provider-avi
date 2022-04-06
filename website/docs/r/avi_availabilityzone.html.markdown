<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_availabilityzone"
sidebar_current: "docs-avi-resource-availabilityzone"
description: |-
  Creates and manages Avi AvailabilityZone.
---

# avi_availabilityzone

The AvailabilityZone resource allows the creation and management of Avi AvailabilityZone

## Example Usage

```hcl
resource "avi_availabilityzone" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Availabilty zone where vcenter list belongs to. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `vcenter_refs` - (Required) Group of vcenter list belong to availabilty zone. It is a reference to an object of type vcenterserver. Field introduced in 20.1.1. Minimum of 1 items required. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `cloud_ref` - (Optional) Availability zone belongs to cloud. It is a reference to an object of type cloud. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Availabilityzone belongs to tenant. It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Availability zone config uuid. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

