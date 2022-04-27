<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_seproperties"
sidebar_current: "docs-avi-datasource-seproperties"
description: |-
  Get information of Avi SeProperties.
---

# avi_seproperties

This data source is used to to get avi_seproperties objects.

## Example Usage

```hcl
data "avi_seproperties" "foo_seproperties" {
    uuid = "seproperties-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SeProperties by name.
* `uuid` - (Optional) Search SeProperties by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `se_agent_properties` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_bootup_properties` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `se_runtime_properties` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

