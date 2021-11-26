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
page_title: "Avi: avi_vsgs"
sidebar_current: "docs-avi-resource-vsgs"
description: |-
  Creates and manages Avi VsGs.
---

# avi_vsgs

The VsGs resource allows the creation and management of Avi VsGs

## Example Usage

```hcl
resource "avi_vsgs" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.3.
* `geodb_uuid` - (Optional) Gslb geodb being associated using this object. Field introduced in 21.1.3.
* `gs_uuid` - (Optional) Gslb service being associated using this object. Field introduced in 21.1.3.
* `gslb_uuid` - (Optional) Gslb being associated using this object. Field introduced in 21.1.3.
* `name` - (Optional) Name of the vs-gs association object. Field introduced in 21.1.3.
* `tenant_ref` - (Optional) Tenant. It is a reference to an object of type tenant. Field introduced in 21.1.3.
* `type` - (Optional) Type of the vs-gs association object. Enum options - VSGS_TYPE_GSLB, VSGS_TYPE_GS, VSGS_TYPE_GEO_DB. Field introduced in 21.1.3.
* `vs_uuid` - (Optional) Virtual service being associated using this object. Field introduced in 21.1.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the vs-gs association object. Field introduced in 21.1.3.

