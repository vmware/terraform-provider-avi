<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_vsgs"
sidebar_current: "docs-avi-datasource-vsgs"
description: |-
  Get information of Avi VsGs.
---

# avi_vsgs

This data source is used to to get avi_vsgs objects.

## Example Usage

```hcl
data "avi_vsgs" "foo_vsgs" {
    uuid = "vsgs-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search VsGs by name.
* `uuid` - (Optional) Search VsGs by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.3. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `geodb_uuid` - Gslb geodb being associated using this object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `gs_uuid` - Gslb service being associated using this object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `gslb_uuid` - Gslb being associated using this object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Name of the vs-gs association object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Tenant. It is a reference to an object of type tenant. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `type` - Type of the vs-gs association object. Enum options - VSGS_TYPE_GSLB, VSGS_TYPE_GS, VSGS_TYPE_GEO_DB. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Uuid of the vs-gs association object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `vs_uuid` - Virtual service being associated using this object. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.

