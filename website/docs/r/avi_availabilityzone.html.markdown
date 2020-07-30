############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `cloud_ref` - (Optional) Availability zone belongs to cloud.
* `name` - (Optional) Availabilty zone where vcenter list belongs to.
* `tenant_ref` - (Optional) Availabilityzone belongs to tenant.
* `vcenter_refs` - (Optional) Group of vcenter list belong to availabilty zone.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Availability zone config uuid.

