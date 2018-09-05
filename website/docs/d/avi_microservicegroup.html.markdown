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
data "MicroServiceGroup" "foo_MicroServiceGroup" {
    uuid = "MicroServiceGroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search MicroServiceGroup by name.
* `uuid` - (Optional) Search MicroServiceGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `created_by` - Creator name.
* `description` - General description.
* `name` - Name of the microservice group.
* `service_refs` - Configure microservice(es).
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the microservice group.

