############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_licenseledgerdetails"
sidebar_current: "docs-avi-datasource-licenseledgerdetails"
description: |-
  Get information of Avi LicenseLedgerDetails.
---

# avi_licenseledgerdetails

This data source is used to to get avi_licenseledgerdetails objects.

## Example Usage

```hcl
data "avi_licenseledgerdetails" "foo_licenseledgerdetails" {
    uuid = "licenseledgerdetails-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search LicenseLedgerDetails by name.
* `uuid` - (Optional) Search LicenseLedgerDetails by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `escrow_infos` - Maintain information about reservation against cookie. Field introduced in 20.1.1.
* `se_infos` - Maintain information about consumed licenses against se_uuid. Field introduced in 20.1.1.
* `tier_usages` - License usage per tier. Field introduced in 20.1.1.
* `uuid` - Uuid for reference. Field introduced in 20.1.1.

