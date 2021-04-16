<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_errorpageprofile"
sidebar_current: "docs-avi-datasource-errorpageprofile"
description: |-
  Get information of Avi ErrorPageProfile.
---

# avi_errorpageprofile

This data source is used to to get avi_errorpageprofile objects.

## Example Usage

```hcl
data "avi_errorpageprofile" "foo_errorpageprofile" {
    uuid = "errorpageprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ErrorPageProfile by name.
* `uuid` - (Optional) Search ErrorPageProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `error_pages` - Defined error pages for http status codes. Field introduced in 17.2.4.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `name` - Field introduced in 17.2.4.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.2.4.
* `uuid` - Field introduced in 17.2.4.

