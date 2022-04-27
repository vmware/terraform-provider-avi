<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_geodb"
sidebar_current: "docs-avi-datasource-geodb"
description: |-
  Get information of Avi GeoDB.
---

# avi_geodb

This data source is used to to get avi_geodb objects.

## Example Usage

```hcl
data "avi_geodb" "foo_geodb" {
    uuid = "geodb-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search GeoDB by name.
* `uuid` - (Optional) Search GeoDB by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - Description. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `files` - Geo database files. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `is_federated` - This field indicates that this object is replicated across gslb federation. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `mappings` - Custom mappings of geo values. All mappings which start with the prefix 'system-' (any case) are reserved for system default objects and may be overwritten. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Geo database name. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Uuid of this object. Field introduced in 21.1.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

