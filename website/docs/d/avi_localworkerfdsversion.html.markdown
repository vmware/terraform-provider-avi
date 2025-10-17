<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_localworkerfdsversion"
sidebar_current: "docs-avi-datasource-localworkerfdsversion"
description: |-
  Get information of Avi LocalWorkerFdsVersion.
---

# avi_localworkerfdsversion

This data source is used to to get avi_localworkerfdsversion objects.

## Example Usage

```hcl
data "avi_localworkerfdsversion" "foo_localworkerfdsversion" {
    uuid = "localworkerfdsversion-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search LocalWorkerFdsVersion by name.
* `uuid` - (Optional) Search LocalWorkerFdsVersion by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Default glw fds version name. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `timeline` - Fds timeline maintained by glw. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Default glw fds version uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `version` - Fds version maintained by glw. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

