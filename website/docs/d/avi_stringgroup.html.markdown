############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_stringgroup"
sidebar_current: "docs-avi-datasource-stringgroup"
description: |-
  Get information of Avi StringGroup.
---

# avi_stringgroup

This data source is used to to get avi_stringgroup objects.

## Example Usage

```hcl
data "avi_stringgroup" "foo_stringgroup" {
    uuid = "stringgroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search StringGroup by name.
* `uuid` - (Optional) Search StringGroup by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - User defined description for the object.
* `kv` - Configure key value in the string group.
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `longest_match` - Enable the longest match, default is the shortest match. Field introduced in 18.2.8.
* `name` - Name of the string group.
* `tenant_ref` - It is a reference to an object of type tenant.
* `type` - Type of stringgroup. Enum options - SG_TYPE_STRING, SG_TYPE_KEYVAL.
* `uuid` - Uuid of the string group.

