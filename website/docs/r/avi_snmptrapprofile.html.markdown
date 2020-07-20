############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_snmptrapprofile"
sidebar_current: "docs-avi-resource-snmptrapprofile"
description: |-
  Creates and manages Avi SnmpTrapProfile.
---

# avi_snmptrapprofile

The SnmpTrapProfile resource allows the creation and management of Avi SnmpTrapProfile

## Example Usage

```hcl
resource "avi_snmptrapprofile" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) A user-friendly name of the snmp trap configuration.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `trap_servers` - (Optional) The ip address or hostname of the snmp trap destination server.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the snmp trap profile object.

