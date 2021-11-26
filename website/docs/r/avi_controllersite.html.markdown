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
page_title: "Avi: avi_controllersite"
sidebar_current: "docs-avi-resource-controllersite"
description: |-
  Creates and manages Avi ControllerSite.
---

# avi_controllersite

The ControllerSite resource allows the creation and management of Avi ControllerSite

## Example Usage

```hcl
resource "avi_controllersite" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `address` - (Required) Ip address or a dns resolvable, fully qualified domain name of the site controller cluster. Field introduced in 18.2.5.
* `name` - (Required) Name for the site controller cluster. Field introduced in 18.2.5.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `port` - (Optional) The controller site cluster's rest api port number. Allowed values are 1-65535. Field introduced in 18.2.5.
* `tenant_ref` - (Optional) Reference for the tenant. It is a reference to an object of type tenant. Field introduced in 18.2.5.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Reference for the site controller cluster. Field introduced in 18.2.5.

