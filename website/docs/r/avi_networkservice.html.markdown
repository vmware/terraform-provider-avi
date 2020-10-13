############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_networkservice"
sidebar_current: "docs-avi-resource-networkservice"
description: |-
  Creates and manages Avi NetworkService.
---

# avi_networkservice

The NetworkService resource allows the creation and management of Avi NetworkService

## Example Usage

```hcl
resource "avi_networkservice" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `labels` - (Optional) Key value pairs for granular object access control.
* `name` - (Optional) Name of the networkservice.
* `routing_service` - (Optional) Routing information of the networkservice.
* `se_group_ref` - (Optional) Service engine group to which the service is applied.
* `service_type` - (Optional) Indicates the type of networkservice.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `vrf_ref` - (Optional) Vrf context to which the service is scoped.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the networkservice.

