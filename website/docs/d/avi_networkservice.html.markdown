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

* `cloud_ref` - It is a reference to an object of type cloud.
* `name` - Name of the networkservice.
* `routing_service` - Routing information of the networkservice.
* `se_group_ref` - Service engine group to which the service is applied.
* `service_type` - Indicates the type of networkservice.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the networkservice.
* `vrf_ref` - Vrf context to which the service is scoped.

