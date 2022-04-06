<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_controllersite"
sidebar_current: "docs-avi-datasource-controllersite"
description: |-
  Get information of Avi ControllerSite.
---

# avi_controllersite

This data source is used to to get avi_controllersite objects.

## Example Usage

```hcl
data "avi_controllersite" "foo_controllersite" {
    uuid = "controllersite-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ControllerSite by name.
* `uuid` - (Optional) Search ControllerSite by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `address` - Ip address or a dns resolvable, fully qualified domain name of the site controller cluster. Field introduced in 18.2.5. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `name` - Name for the site controller cluster. Field introduced in 18.2.5. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `port` - The controller site cluster's rest api port number. Allowed values are 1-65535. Field introduced in 18.2.5. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - Reference for the tenant. It is a reference to an object of type tenant. Field introduced in 18.2.5. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Reference for the site controller cluster. Field introduced in 18.2.5. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

