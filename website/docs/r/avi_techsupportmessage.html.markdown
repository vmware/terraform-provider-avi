<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_techsupportmessage"
sidebar_current: "docs-avi-resource-techsupportmessage"
description: |-
  Creates and manages Avi TechSupportMessage.
---

# avi_techsupportmessage

The TechSupportMessage resource allows the creation and management of Avi TechSupportMessage

## Example Usage

```hcl
resource "avi_techsupportmessage" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `status` - (Optional) 'techsupport status for the current invocation.'. Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `status_code` - (Optional) 'techsupport status code for the current invocation.'. Enum options - SYSERR_SUCCESS, SYSERR_FAILURE, SYSERR_OUT_OF_MEMORY, SYSERR_NO_ENT, SYSERR_INVAL, SYSERR_ACCESS, SYSERR_FAULT, SYSERR_IO, SYSERR_TIMEOUT, SYSERR_NOT_SUPPORTED, SYSERR_NOT_READY, SYSERR_UPGRADE_IN_PROGRESS, SYSERR_WARM_START_IN_PROGRESS, SYSERR_TRY_AGAIN, SYSERR_NOT_UPGRADING, SYSERR_PENDING, SYSERR_EVENT_GEN_FAILURE, SYSERR_CONFIG_PARAM_MISSING, SYSERR_RANGE, SYSERR_FAILED... Field introduced in 18.2.3. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition. 
* `tech_support_ref` - (Optional) 'techsupport object ref.'. It is a reference to an object of type techsupport. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition. 


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:


