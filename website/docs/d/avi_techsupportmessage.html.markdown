<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_techsupportmessage"
sidebar_current: "docs-avi-datasource-techsupportmessage"
description: |-
  Get information of Avi TechSupportMessage.
---

# avi_techsupportmessage

This data source is used to to get avi_techsupportmessage objects.

## Example Usage

```hcl
data "avi_techsupportmessage" "foo_techsupportmessage" {
    uuid = "techsupportmessage-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TechSupportMessage by name.
* `uuid` - (Optional) Search TechSupportMessage by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `status` - 'techsupport status for the current invocation.'. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `status_code` - 'techsupport status code for the current invocation.'. Enum options - SYSERR_SUCCESS, SYSERR_FAILURE, SYSERR_OUT_OF_MEMORY, SYSERR_NO_ENT, SYSERR_INVAL, SYSERR_ACCESS, SYSERR_FAULT, SYSERR_IO, SYSERR_TIMEOUT, SYSERR_NOT_SUPPORTED, SYSERR_NOT_READY, SYSERR_UPGRADE_IN_PROGRESS, SYSERR_WARM_START_IN_PROGRESS, SYSERR_TRY_AGAIN, SYSERR_NOT_UPGRADING, SYSERR_PENDING, SYSERR_EVENT_GEN_FAILURE, SYSERR_CONFIG_PARAM_MISSING, SYSERR_RANGE, SYSERR_FAILED... Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tech_support_ref` - 'techsupport object ref.'. It is a reference to an object of type techsupport. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

