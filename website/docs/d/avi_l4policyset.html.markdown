<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_l4policyset"
sidebar_current: "docs-avi-datasource-l4policyset"
description: |-
  Get information of Avi L4PolicySet.
---

# avi_l4policyset

This data source is used to to get avi_l4policyset objects.

## Example Usage

```hcl
data "avi_l4policyset" "foo_l4policyset" {
    uuid = "l4policyset-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search L4PolicySet by name.
* `uuid` - (Optional) Search L4PolicySet by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `created_by` - Creator name. Field introduced in 17.2.7.
* `description` - Field introduced in 17.2.7.
* `is_internal_policy` - Field introduced in 17.2.7.
* `l4_connection_policy` - Policy to apply when a new transport connection is setup. Field introduced in 17.2.7.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `name` - Name of the l4 policy set. Field introduced in 17.2.7.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.2.7.
* `uuid` - Id of the l4 policy set. Field introduced in 17.2.7.

