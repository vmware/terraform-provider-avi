<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_alertemailconfig"
sidebar_current: "docs-avi-datasource-alertemailconfig"
description: |-
  Get information of Avi AlertEmailConfig.
---

# avi_alertemailconfig

This data source is used to to get avi_alertemailconfig objects.

## Example Usage

```hcl
data "avi_alertemailconfig" "foo_alertemailconfig" {
    uuid = "alertemailconfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AlertEmailConfig by name.
* `uuid` - (Optional) Search AlertEmailConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cc_emails` - Alerts are copied to the comma separated list of  email recipients. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `description` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `name` - A user-friendly name of the email notification service. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `to_emails` - Alerts are sent to the comma separated list of  email recipients. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

