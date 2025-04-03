<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_gslbsmruntime"
sidebar_current: "docs-avi-datasource-gslbsmruntime"
description: |-
  Get information of Avi GslbSMRuntime.
---

# avi_gslbsmruntime

This data source is used to to get avi_gslbsmruntime objects.

## Example Usage

```hcl
data "avi_gslbsmruntime" "foo_gslbsmruntime" {
    uuid = "gslbsmruntime-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search GslbSMRuntime by name.
* `uuid` - (Optional) Search GslbSMRuntime by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cluster_leader` - The controller cluster leader node uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `cluster_uuid` - The site controller cluster uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `dns_configs` - Sub domain configuration for the gslb. Gslb service's fqdn must be a match one of these subdomains. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `dns_info` - Dns info at the site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enabled` - Activate/de-activate state retrieved from the cfg. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `events` - Captures sm related events. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `health_monitor_info` - This field will provide information on origin(site name) of the health monitoring information. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `leader_cluster_uuid` - Mark this site as leader of gslb configuration. This site is the one among the avi sites. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `member_type` - The site's member type  a leader is set to active while all members are set to passive. Enum options - GSLB_ACTIVE_MEMBER, GSLB_PASSIVE_MEMBER. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - The name of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `node_uuid` - The controller cluster node uuid that processes the site.sites are sharded across the cluster nodes. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `num_of_retries` - Number of retry attempts to reach the remote site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_uuid` - Gslb sm runtime object uuid. Points to the gslb to which this belongs. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `oper_status` - Gslb site operational status, represents whether site is up or down. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `remote_info` - Remote info is basically updated by grw. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `role` - Site role  leader or follower. Enum options - GSLB_LEADER, GSLB_MEMBER, GSLB_NOT_A_MEMBER. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `site_name` - The gslb site name. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `site_type` - Indicates if it is avi site or third-party. Enum options - GSLB_AVI_SITE, GSLB_THIRD_PARTY_SITE. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `state` - Represents the state of the site. Enum options - SITE_STATE_NULL, SITE_STATE_JOIN_IN_PROGRESS, SITE_STATE_LEAVE_IN_PROGRESS, SITE_STATE_INIT, SITE_STATE_UNREACHABLE, SITE_STATE_MMODE, SITE_STATE_DISABLE_IN_PROGRESS, SITE_STATE_DISABLED, SITE_STATE_HS_IN_PROGRESS. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `sw_version` - Current software version of the site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - The uuid of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `view_id` - The view-id is used in change-leader mode to differentiate partitioned groups while they have the same gslb namespace. Each partitioned group will be able to operate independently by using the view-id. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

