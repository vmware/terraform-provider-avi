<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_alertemailconfig"
sidebar_current: "docs-avi-resource-alertemailconfig"
description: |-
  Creates and manages Avi AlertEmailConfig.
---

# avi_alertemailconfig

The AlertEmailConfig resource allows the creation and management of Avi AlertEmailConfig

## Example Usage

```hcl
resource "avi_alertemailconfig" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user-friendly name of the email notification service. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `to_emails` - (Required) Alerts are sent to the comma separated list of  email recipients. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `cc_emails` - (Optional) Alerts are copied to the comma separated list of  email recipients. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `description` - (Optional) Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

