---
layout: "avi"
page_title: "AVI: avi_availabilityzone"
sidebar_current: "docs-avi-datasource-availabilityzone"
description: |-
  Get information of Avi AvailabilityZone.
---

# avi_availabilityzone

This data source is used to to get avi_availabilityzone objects.

## Example Usage

```hcl
data "avi_availabilityzone" "foo_availabilityzone" {
    uuid = "availabilityzone-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search AvailabilityZone by name.
* `uuid` - (Optional) Search AvailabilityZone by uuid.
* `cloud_ref` - (Optional) Search AvailabilityZone by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_ref` - Availability zone belongs to cloud.
* `name` - Availabilty zone where vcenter list belongs to.
* `tenant_ref` - Availabilityzone belongs to tenant.
* `uuid` - Availability zone config uuid.
* `vcenter_refs` - Group of vcenter list belong to availabilty zone.

