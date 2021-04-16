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

* `created_by` - Creator name.
* `description` - User defined description for the object.
* `name` - Name of the microservice group.
* `service_refs` - Configure microservice(es). It is a reference to an object of type microservice.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the microservice group.

