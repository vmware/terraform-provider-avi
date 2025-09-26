<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_gslbcrmruntime"
sidebar_current: "docs-avi-datasource-gslbcrmruntime"
description: |-
  Get information of Avi GslbCRMRuntime.
---

# avi_gslbcrmruntime

This data source is used to to get avi_gslbcrmruntime objects.

## Example Usage

```hcl
data "avi_gslbcrmruntime" "foo_gslbcrmruntime" {
    uuid = "gslbcrmruntime-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search GslbCRMRuntime by name.
* `uuid` - (Optional) Search GslbCRMRuntime by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cluster_uuid` - This field tracks the site_uuid for local/remote site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `events` - Events captured wrt to config replication. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `fds_info` - Federated data store related info. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `local_info` - Represents local info for the site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - The name of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_uuid` - Gslb crm runtime object uuid. Points to the gslb to which this belongs. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `remote_info` - Respresents remote site's info wrt to replication. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `replication_policy` - Policy for replicating configuration to the active follower sites. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `site_name` - This field tracks the site name. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `status_info` - Crm operational status. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - The uuid of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

