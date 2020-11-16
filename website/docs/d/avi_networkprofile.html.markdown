############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_networkprofile"
sidebar_current: "docs-avi-datasource-networkprofile"
description: |-
  Get information of Avi NetworkProfile.
---

# avi_networkprofile

This data source is used to to get avi_networkprofile objects.

## Example Usage

```hcl
data "avi_networkprofile" "foo_networkprofile" {
    uuid = "networkprofile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search NetworkProfile by name.
* `uuid` - (Optional) Search NetworkProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `connection_mirror` - When enabled, avi mirrors all tcp fastpath connections to standby.
* `description` - User defined description for the object.
* `labels` - Key value pairs for granular object access control.
* `name` - The name of the network profile.
* `profile` - Dict settings for networkprofile.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the network profile.

