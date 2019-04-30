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
data "SnmpTrapProfile" "foo_SnmpTrapProfile" {
    uuid = "SnmpTrapProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search SnmpTrapProfile by name.
* `uuid` - (Optional) Search SnmpTrapProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - A user-friendly name of the snmp trap configuration.
* `tenant_ref` - It is a reference to an object of type tenant.
* `trap_servers` - The ip address or hostname of the snmp trap destination server.
* `uuid` - Uuid of the snmp trap profile object.

