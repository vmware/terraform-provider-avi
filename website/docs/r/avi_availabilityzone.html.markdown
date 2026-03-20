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

* `name` - (Required) Availabilty zone where vcenter list belongs to. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. Changing this value forces the resource to be recreated.
* `az_clusters` - (Optional) Group of clusters belongs to the az. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `az_datastores` - (Optional) Group of datastores associated with the az. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `az_hosts` - (Optional) Group of hosts associated with the az. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 
* `cloud_ref` - (Optional) Availability zone belongs to cloud. It is a reference to an object of type cloud. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. Changing this value forces the resource to be recreated.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tenant_ref` - (Optional) Availabilityzone belongs to tenant. It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `vsphere_zones` - (Optional) Vsphere zone associated with the az. Field introduced in 32.1.1. Maximum of 1 items allowed. Allowed with any value in enterprise, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Availability zone config uuid. Field introduced in 20.1.1. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

