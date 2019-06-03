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

* `address` - Ip address or a dns resolvable, fully qualified domain name of the site controller cluster.
* `name` - Name for the site controller cluster.
* `port` - The controller site cluster's rest api port number.
* `tenant_ref` - Reference for the tenant.
* `uuid` - Reference for the site controller cluster.

