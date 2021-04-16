<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_wafapplicationsignatureprovider"
sidebar_current: "docs-avi-datasource-wafapplicationsignatureprovider"
description: |-
  Get information of Avi WafApplicationSignatureProvider.
---

# avi_wafapplicationsignatureprovider

This data source is used to to get avi_wafapplicationsignatureprovider objects.

## Example Usage

```hcl
data "avi_wafapplicationsignatureprovider" "foo_wafapplicationsignatureprovider" {
    uuid = "wafapplicationsignatureprovider-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search WafApplicationSignatureProvider by name.
* `uuid` - (Optional) Search WafApplicationSignatureProvider by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of application specific ruleset provider. Field introduced in 20.1.1.
* `service_status` - If this object is managed by the application signatures update  service, this field contain the status of this syncronization. Field introduced in 20.1.3.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 20.1.1.
* `uuid` - Field introduced in 20.1.1.

