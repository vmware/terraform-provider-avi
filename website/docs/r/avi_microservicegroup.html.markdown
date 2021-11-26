############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_microservicegroup"
sidebar_current: "docs-avi-resource-microservicegroup"
description: |-
  Creates and manages Avi MicroServiceGroup.
---

# avi_microservicegroup

The MicroServiceGroup resource allows the creation and management of Avi MicroServiceGroup

## Example Usage

```hcl
resource "avi_microservicegroup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the microservice group.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `created_by` - (Optional) Creator name.
* `description` - (Optional) User defined description for the object.
* `service_refs` - (Optional) Configure microservice(es). It is a reference to an object of type microservice.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the microservice group.

