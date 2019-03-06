
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
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
data "VSDataScriptSet" "foo_VSDataScriptSet" {
    uuid = "VSDataScriptSet-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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
* `description` - General description.
* `ipgroup_refs` - Uuid of ip groups that could be referred by vsdatascriptset objects.
* `name` - Name for the virtual service datascript collection.
* `pool_group_refs` - Uuid of pool groups that could be referred by vsdatascriptset objects.
* `pool_refs` - Uuid of pools that could be referred by vsdatascriptset objects.
* `string_group_refs` - Uuid of string groups that could be referred by vsdatascriptset objects.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the virtual service datascript collection.

