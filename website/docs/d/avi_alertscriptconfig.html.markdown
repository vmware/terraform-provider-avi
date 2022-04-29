<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_alertscriptconfig"
sidebar_current: "docs-avi-datasource-alertscriptconfig"
description: |-
  Get information of Avi AlertScriptConfig.
---

# avi_alertscriptconfig

This data source is used to to get avi_alertscriptconfig objects.

## Example Usage

```hcl
data "avi_alertscriptconfig" "foo_alertscriptconfig" {
    uuid = "alertscriptconfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AlertScriptConfig by name.
* `uuid` - (Optional) Search AlertScriptConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `action_script` - User defined alert action script. Please refer to kb.avinetworks.com for more information. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - A user-friendly name of the script. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

