<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_memorybalancerrequest"
sidebar_current: "docs-avi-datasource-memorybalancerrequest"
description: |-
  Get information of Avi MemoryBalancerRequest.
---

# avi_memorybalancerrequest

This data source is used to to get avi_memorybalancerrequest objects.

## Example Usage

```hcl
data "avi_memorybalancerrequest" "foo_memorybalancerrequest" {
    uuid = "memorybalancerrequest-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search MemoryBalancerRequest by name.
* `uuid` - (Optional) Search MemoryBalancerRequest by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `controller_info` - Current details regarding controller. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `name` - Name of controller process. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `node_uuid` - Uuid of node. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `process_info` - Current process information of the controller process. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `process_instance` - Instance of the controller process. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `tenant_ref` - Uuid of tenant object. It is a reference to an object of type tenant. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `timestamp` - Time at which memory balancer request was created/updated. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `uuid` - Uuid of memory balancer request object. Field introduced in 21.1.1. Allowed in enterprise with any value edition, enterprise with cloud services edition.

