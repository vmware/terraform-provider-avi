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
page_title: "AVI: avi_controllerportalregistration"
sidebar_current: "docs-avi-datasource-controllerportalregistration"
description: |-
  Get information of Avi ControllerPortalRegistration.
---

# avi_controllerportalregistration

This data source is used to to get avi_controllerportalregistration objects.

## Example Usage

```hcl
data "avi_controllerportalregistration" "foo_controllerportalregistration" {
    uuid = "controllerportalregistration-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ControllerPortalRegistration by name.
* `uuid` - (Optional) Search ControllerPortalRegistration by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `asset` - Field introduced in 18.2.6.
* `name` - Field introduced in 18.2.6.
* `portal_auth` - Field introduced in 18.2.6.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 18.2.6.
* `uuid` - Field introduced in 18.2.6.

