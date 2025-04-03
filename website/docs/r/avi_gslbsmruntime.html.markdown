<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_gslbsmruntime"
sidebar_current: "docs-avi-resource-gslbsmruntime"
description: |-
  Creates and manages Avi GslbSMRuntime.
---

# avi_gslbsmruntime

The GslbSMRuntime resource allows the creation and management of Avi GslbSMRuntime

## Example Usage

```hcl
resource "avi_gslbsmruntime" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `leader_cluster_uuid` - (Required) Mark this site as leader of gslb configuration. This site is the one among the avi sites. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `cluster_leader` - (Optional) The controller cluster leader node uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `cluster_uuid` - (Optional) The site controller cluster uuid. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `dns_configs` - (Optional) Sub domain configuration for the gslb. Gslb service's fqdn must be a match one of these subdomains. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `dns_info` - (Optional) Dns info at the site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `enabled` - (Optional) Activate/de-activate state retrieved from the cfg. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `events` - (Optional) Captures sm related events. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `health_monitor_info` - (Optional) This field will provide information on origin(site name) of the health monitoring information. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `member_type` - (Optional) The site's member type  a leader is set to active while all members are set to passive. Enum options - GSLB_ACTIVE_MEMBER, GSLB_PASSIVE_MEMBER. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - (Optional) The name of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `node_uuid` - (Optional) The controller cluster node uuid that processes the site.sites are sharded across the cluster nodes. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `num_of_retries` - (Optional) Number of retry attempts to reach the remote site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_uuid` - (Optional) Gslb sm runtime object uuid. Points to the gslb to which this belongs. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `oper_status` - (Optional) Gslb site operational status, represents whether site is up or down. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `remote_info` - (Optional) Remote info is basically updated by grw. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `role` - (Optional) Site role  leader or follower. Enum options - GSLB_LEADER, GSLB_MEMBER, GSLB_NOT_A_MEMBER. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `site_name` - (Optional) The gslb site name. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `site_type` - (Optional) Indicates if it is avi site or third-party. Enum options - GSLB_AVI_SITE, GSLB_THIRD_PARTY_SITE. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `state` - (Optional) Represents the state of the site. Enum options - SITE_STATE_NULL, SITE_STATE_JOIN_IN_PROGRESS, SITE_STATE_LEAVE_IN_PROGRESS, SITE_STATE_INIT, SITE_STATE_UNREACHABLE, SITE_STATE_MMODE, SITE_STATE_DISABLE_IN_PROGRESS, SITE_STATE_DISABLED, SITE_STATE_HS_IN_PROGRESS. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `sw_version` - (Optional) Current software version of the site. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Uuid of the tenant. It is a reference to an object of type tenant. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `view_id` - (Optional) The view-id is used in change-leader mode to differentiate partitioned groups while they have the same gslb namespace. Each partitioned group will be able to operate independently by using the view-id. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  The uuid of db entry. Field introduced in 31.1.1. Allowed with any value in enterprise, enterprise with cloud services edition.

