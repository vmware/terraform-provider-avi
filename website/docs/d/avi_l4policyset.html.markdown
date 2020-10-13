############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `created_by` - Creator name.
* `description` - Field introduced in 17.2.7.
* `is_internal_policy` - Field introduced in 17.2.7.
* `l4_connection_policy` - Policy to apply when a new transport connection is setup.
* `labels` - Key value pairs for granular object access control.
* `name` - Name of the l4 policy set.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Id of the l4 policy set.

