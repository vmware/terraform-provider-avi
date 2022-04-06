<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_microservicegroup"
sidebar_current: "docs-avi-datasource-microservicegroup"
description: |-
  Get information of Avi MicroServiceGroup.
---

# avi_microservicegroup

This data source is used to to get avi_microservicegroup objects.

## Example Usage

```hcl
data "avi_microservicegroup" "foo_microservicegroup" {
    uuid = "microservicegroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search MicroServiceGroup by name.
* `uuid` - (Optional) Search MicroServiceGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `created_by` - Creator name. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `description` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `name` - Name of the microservice group. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `service_refs` - Configure microservice(es). It is a reference to an object of type microservice. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Uuid of the microservice group. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

