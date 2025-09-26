<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_localworkerfdsversion"
sidebar_current: "docs-avi-resource-localworkerfdsversion"
description: |-
  Creates and manages Avi LocalWorkerFdsVersion.
---

# avi_localworkerfdsversion

The LocalWorkerFdsVersion resource allows the creation and management of Avi LocalWorkerFdsVersion

## Example Usage

```hcl
resource "avi_localworkerfdsversion" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional) Default glw fds version name. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `timeline` - (Optional) Fds timeline maintained by glw. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `version` - (Optional) Fds version maintained by glw. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Default glw fds version uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

