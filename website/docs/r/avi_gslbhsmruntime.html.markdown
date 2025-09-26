<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_gslbhsmruntime"
sidebar_current: "docs-avi-resource-gslbhsmruntime"
description: |-
  Creates and manages Avi GslbHSMRuntime.
---

# avi_gslbhsmruntime

The GslbHSMRuntime resource allows the creation and management of Avi GslbHSMRuntime

## Example Usage

```hcl
resource "avi_gslbhsmruntime" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `cluster_uuid` - (Optional) The site controller cluster uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enabled` - (Optional) Represents whether hsm is enabled/disabled. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `events` - (Optional) Events captured wrt to config replication. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `local_info` - (Optional) Represents local info for the site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - (Optional) The name of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_uuid` - (Optional) Gslb hsm runtime object uuid. Points to the gslb to which this belongs. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `oper_status` - (Optional) Gslb site operational status, represents whether site is up or down. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `remote_info` - (Optional) Remote info is basically updated by grw. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `send_interval` - (Optional) Frequency with which group members communicate. This field shadows glb_cfg.send_interval. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `site_name` - (Optional) The gslb site name. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  The uuid of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

