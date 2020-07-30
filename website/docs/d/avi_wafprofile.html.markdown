############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_wafprofile"
sidebar_current: "docs-avi-datasource-wafprofile"
description: |-
  Get information of Avi WafProfile.
---

# avi_wafprofile

This data source is used to to get avi_wafprofile objects.

## Example Usage

```hcl
data "avi_wafprofile" "foo_wafprofile" {
    uuid = "wafprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafProfile by name.
* `uuid` - (Optional) Search WafProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `config` - Config params for waf.
* `description` - Field introduced in 17.2.1.
* `files` - List of data files used for waf rules.
* `name` - Field introduced in 17.2.1.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Field introduced in 17.2.1.

