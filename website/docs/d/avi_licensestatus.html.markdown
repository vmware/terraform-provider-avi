############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_licensestatus"
sidebar_current: "docs-avi-datasource-licensestatus"
description: |-
  Get information of Avi LicenseStatus.
---

# avi_licensestatus

This data source is used to to get avi_licensestatus objects.

## Example Usage

```hcl
data "avi_licensestatus" "foo_licensestatus" {
    uuid = "licensestatus-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search LicenseStatus by name.
* `uuid` - (Optional) Search LicenseStatus by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.3.
* `name` - Name. Field introduced in 21.1.3.
* `saas_status` - Saas licensing status. Field introduced in 21.1.3.
* `tenant_ref` - Tenant. It is a reference to an object of type tenant. Field introduced in 21.1.3.
* `uuid` - Uuid. Field introduced in 21.1.3.

