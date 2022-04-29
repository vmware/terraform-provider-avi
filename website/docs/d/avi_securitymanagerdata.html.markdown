<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_securitymanagerdata"
sidebar_current: "docs-avi-datasource-securitymanagerdata"
description: |-
  Get information of Avi SecurityManagerData.
---

# avi_securitymanagerdata

This data source is used to to get avi_securitymanagerdata objects.

## Example Usage

```hcl
data "avi_securitymanagerdata" "foo_securitymanagerdata" {
    uuid = "securitymanagerdata-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SecurityManagerData by name.
* `uuid` - (Optional) Search SecurityManagerData by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `app_learning_info` - Information about various applications. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Virtualservice name. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Virtualservice uuid. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

