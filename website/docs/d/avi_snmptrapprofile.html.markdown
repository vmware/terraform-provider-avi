<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_snmptrapprofile"
sidebar_current: "docs-avi-datasource-snmptrapprofile"
description: |-
  Get information of Avi SnmpTrapProfile.
---

# avi_snmptrapprofile

This data source is used to to get avi_snmptrapprofile objects.

## Example Usage

```hcl
data "avi_snmptrapprofile" "foo_snmptrapprofile" {
    uuid = "snmptrapprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SnmpTrapProfile by name.
* `uuid` - (Optional) Search SnmpTrapProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - A user-friendly name of the snmp trap configuration. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `trap_servers` - The ip address or hostname of the snmp trap destination server. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the snmp trap profile object. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

