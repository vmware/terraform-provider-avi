
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
page_title: "AVI: avi_l4policyset"
sidebar_current: "docs-avi-datasource-l4policyset"
description: |-
  Get information of Avi L4PolicySet.
---

# avi_l4policyset

This data source is used to to get avi_l4policyset objects.

## Example Usage

```hcl
data "L4PolicySet" "foo_L4PolicySet" {
    uuid = "L4PolicySet-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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
* `name` - Name of the l4 policy set.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Id of the l4 policy set.

