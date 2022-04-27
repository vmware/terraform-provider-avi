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

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.3. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `saas_status` - Saas licensing status. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `service_update` - Pulse license service update. Field introduced in 21.1.4. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Uuid. Field introduced in 21.1.3. Allowed in enterprise edition with any value, enterprise with cloud services edition.

