############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_customerportalinfo"
sidebar_current: "docs-avi-datasource-customerportalinfo"
description: |-
  Get information of Avi CustomerPortalInfo.
---

# avi_customerportalinfo

This data source is used to to get avi_customerportalinfo objects.

## Example Usage

```hcl
data "avi_customerportalinfo" "foo_customerportalinfo" {
    uuid = "customerportalinfo-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CustomerPortalInfo by name.
* `uuid` - (Optional) Search CustomerPortalInfo by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `polling_interval` - Time interval in minutes.
* `portal_url` - The fqdn or ip address of the customer portal.
* `uuid` - Field introduced in 18.2.6.

