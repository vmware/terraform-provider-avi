
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
page_title: "AVI: avi_tenant"
sidebar_current: "docs-avi-datasource-tenant"
description: |-
  Get information of Avi Tenant.
---

# avi_tenant

This data source is used to to get avi_tenant objects.

## Example Usage

```hcl
data "Tenant" "foo_Tenant" {
    uuid = "Tenant-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Tenant by name.
* `uuid` - (Optional) Search Tenant by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `config_settings` - General description.
* `created_by` - Creator of this tenant.
* `description` - General description.
* `local` - General description.
* `name` - General description.
* `uuid` - General description.

