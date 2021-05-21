<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_errorpagebody"
sidebar_current: "docs-avi-datasource-errorpagebody"
description: |-
  Get information of Avi ErrorPageBody.
---

# avi_errorpagebody

This data source is used to to get avi_errorpagebody objects.

## Example Usage

```hcl
data "avi_errorpagebody" "foo_errorpagebody" {
    uuid = "errorpagebody-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ErrorPageBody by name.
* `uuid` - (Optional) Search ErrorPageBody by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `error_page_body` - Error page body sent to client when match. Field introduced in 17.2.4.
* `format` - Format of an error page body html or json. Enum options - ERROR_PAGE_FORMAT_HTML, ERROR_PAGE_FORMAT_JSON. Field introduced in 18.2.3.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `name` - Field introduced in 17.2.4.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.2.4.
* `uuid` - Field introduced in 17.2.4.

