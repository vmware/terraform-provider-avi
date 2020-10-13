############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_vsdatascriptset"
sidebar_current: "docs-avi-datasource-vsdatascriptset"
description: |-
  Get information of Avi VSDataScriptSet.
---

# avi_vsdatascriptset

This data source is used to to get avi_vsdatascriptset objects.

## Example Usage

```hcl
data "avi_vsdatascriptset" "foo_vsdatascriptset" {
    uuid = "vsdatascriptset-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search VSDataScriptSet by name.
* `uuid` - (Optional) Search VSDataScriptSet by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `created_by` - Creator name.
* `datascript` - Datascripts to execute.
* `description` - User defined description for the object.
* `ipgroup_refs` - Uuid of ip groups that could be referred by vsdatascriptset objects.
* `labels` - Key value pairs for granular object access control.
* `name` - Name for the virtual service datascript collection.
* `pool_group_refs` - Uuid of pool groups that could be referred by vsdatascriptset objects.
* `pool_refs` - Uuid of pools that could be referred by vsdatascriptset objects.
* `protocol_parser_refs` - List of protocol parsers that could be referred by vsdatascriptset objects.
* `rate_limiters` - The rate limit definitions needed for this datascript.
* `string_group_refs` - Uuid of string groups that could be referred by vsdatascriptset objects.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the virtual service datascript collection.

