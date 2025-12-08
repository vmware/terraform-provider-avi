<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_gslbhsmruntime"
sidebar_current: "docs-avi-datasource-gslbhsmruntime"
description: |-
  Get information of Avi GslbHSMRuntime.
---

# avi_gslbhsmruntime

This data source is used to to get avi_gslbhsmruntime objects.

## Example Usage

```hcl
data "avi_gslbhsmruntime" "foo_gslbhsmruntime" {
    uuid = "gslbhsmruntime-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search GslbHSMRuntime by name.
* `uuid` - (Optional) Search GslbHSMRuntime by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cluster_uuid` - The site controller cluster uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enabled` - Represents whether hsm is enabled/disabled. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `events` - Events captured wrt to config replication. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `local_info` - Represents local info for the site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - The name of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_uuid` - Gslb hsm runtime object uuid. Points to the gslb to which this belongs. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `oper_status` - Gslb site operational status, represents whether site is up or down. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `remote_info` - Remote info is basically updated by grw. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `send_interval` - Frequency with which group members communicate. This field shadows glb_cfg.send_interval. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `site_name` - The gslb site name. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - The uuid of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

