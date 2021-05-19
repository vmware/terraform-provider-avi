<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_clusterclouddetails"
sidebar_current: "docs-avi-resource-clusterclouddetails"
description: |-
  Creates and manages Avi ClusterCloudDetails.
---

# avi_clusterclouddetails

The ClusterCloudDetails resource allows the creation and management of Avi ClusterCloudDetails

## Example Usage

```hcl
resource "avi_clusterclouddetails" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Field introduced in 17.2.5.
* `azure_info` - (Optional) Azure info to configure cluster_vip on the controller. Field introduced in 17.2.5.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 17.2.5.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Field introduced in 17.2.5.

