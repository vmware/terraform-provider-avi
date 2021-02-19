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

* `error_page_body` - Error page body sent to client when match. Field introduced in 17.2.4.
* `format` - Format of an error page body html or json. Enum options - ERROR_PAGE_FORMAT_HTML, ERROR_PAGE_FORMAT_JSON. Field introduced in 18.2.3.
* `labels` - Key value pairs for granular object access control. Also allows for classification and tagging of similar objects. Field introduced in 20.1.2. Maximum of 4 items allowed.
* `name` - Field introduced in 17.2.4.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.2.4.
* `uuid` - Field introduced in 17.2.4.

