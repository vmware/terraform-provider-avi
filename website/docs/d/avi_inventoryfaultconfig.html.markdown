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
page_title: "AVI: avi_inventoryfaultconfig"
sidebar_current: "docs-avi-datasource-inventoryfaultconfig"
description: |-
  Get information of Avi InventoryFaultConfig.
---

# avi_inventoryfaultconfig

This data source is used to to get avi_inventoryfaultconfig objects.

## Example Usage

```hcl
data "avi_inventoryfaultconfig" "foo_inventoryfaultconfig" {
    uuid = "inventoryfaultconfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search InventoryFaultConfig by name.
* `uuid` - (Optional) Search InventoryFaultConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `controller_faults` - Configure controller faults. Field introduced in 20.1.6.
* `name` - Name. Field introduced in 20.1.6.
* `serviceengine_faults` - Configure serviceengine faults. Field introduced in 20.1.6.
* `tenant_ref` - Tenant. It is a reference to an object of type tenant. Field introduced in 20.1.6.
* `uuid` - Uuid auto generated. Field introduced in 20.1.6.
* `virtualservice_faults` - Configure virtualservice faults. Field introduced in 20.1.6.

