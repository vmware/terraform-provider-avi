<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_networkservice"
sidebar_current: "docs-avi-datasource-networkservice"
description: |-
  Get information of Avi NetworkService.
---

# avi_networkservice

This data source is used to to get avi_networkservice objects.

## Example Usage

```hcl
data "avi_networkservice" "foo_networkservice" {
    uuid = "networkservice-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search NetworkService by name.
* `uuid` - (Optional) Search NetworkService by uuid.
* `cloud_ref` - (Optional) Search NetworkService by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_ref` - It is a reference to an object of type cloud. Field introduced in 18.2.5.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `name` - Name of the networkservice. Field introduced in 18.2.5.
* `routing_service` - Routing information of the networkservice. Field introduced in 18.2.5.
* `se_group_ref` - Service engine group to which the service is applied. It is a reference to an object of type serviceenginegroup. Field introduced in 18.2.5.
* `service_type` - Indicates the type of networkservice. Enum options - ROUTING_SERVICE. Field introduced in 18.2.5.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 18.2.5.
* `uuid` - Uuid of the networkservice. Field introduced in 18.2.5.
* `vrf_ref` - Vrf context to which the service is scoped. It is a reference to an object of type vrfcontext. Field introduced in 18.2.5.

