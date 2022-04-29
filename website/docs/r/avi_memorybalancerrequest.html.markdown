<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_memorybalancerrequest"
sidebar_current: "docs-avi-resource-memorybalancerrequest"
description: |-
  Creates and manages Avi MemoryBalancerRequest.
---

# avi_memorybalancerrequest

The MemoryBalancerRequest resource allows the creation and management of Avi MemoryBalancerRequest

## Example Usage

```hcl
resource "avi_memorybalancerrequest" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of controller process. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `controller_info` - (Optional) Current details regarding controller. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `node_uuid` - (Optional) Uuid of node. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `process_info` - (Optional) Current process information of the controller process. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `process_instance` - (Optional) Instance of the controller process. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Uuid of tenant object. It is a reference to an object of type tenant. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `timestamp` - (Optional) Time at which memory balancer request was created/updated. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of memory balancer request object. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

