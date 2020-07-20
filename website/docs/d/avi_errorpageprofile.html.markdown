############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `error_pages` - Defined error pages for http status codes.
* `name` - Field introduced in 17.2.4.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Field introduced in 17.2.4.

