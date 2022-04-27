<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `connection_mirror` - When enabled, avi mirrors all tcp fastpath connections to standby. Applicable only in legacy ha mode. Field introduced in 18.1.3,18.2.1. Allowed in enterprise edition with any value, essentials edition(allowed values- false), basic, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - The name of the network profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `profile` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the network profile. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

