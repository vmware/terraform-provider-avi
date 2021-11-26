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
page_title: "Avi: avi_securitymanagerdata"
sidebar_current: "docs-avi-resource-securitymanagerdata"
description: |-
  Creates and manages Avi SecurityManagerData.
---

# avi_securitymanagerdata

The SecurityManagerData resource allows the creation and management of Avi SecurityManagerData

## Example Usage

```hcl
resource "avi_securitymanagerdata" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Virtualservice name. Field introduced in 20.1.1.
* `app_learning_info` - (Optional) Information about various applications. Field introduced in 20.1.1.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Field introduced in 20.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Virtualservice uuid. Field introduced in 20.1.1.

