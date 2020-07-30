############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `error_page_body` - Error page body sent to client when match.
* `format` - Format of an error page body html or json.
* `name` - Field introduced in 17.2.4.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Field introduced in 17.2.4.

