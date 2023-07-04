<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_gslb"
sidebar_current: "docs-avi-datasource-gslb"
description: |-
  Get information of Avi Gslb.
---

# avi_gslb

This data source is used to to get avi_gslb objects.

## Example Usage

```hcl
data "avi_gslb" "foo_gslb" {
    uuid = "gslb-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Gslb by name.
* `uuid` - (Optional) Search Gslb by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `async_interval` - Frequency with which messages are propagated to vs mgr. Value of 0 disables async behavior and rpc are sent inline. Allowed values are 0-5. Field introduced in 18.2.3. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `clear_on_max_retries` - Max retries after which the remote site is treated as a fresh start. In fresh start all the configs are downloaded. Allowed values are 1-1024. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `client_ip_addr_group` - Group to specify if the client ip addresses are public or private. Field introduced in 17.1.2. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `dns_configs` - Sub domain configuration for the gslb. Gslb service's fqdn must be a match one of these subdomains. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `enable_config_by_members` - Allows enable/disable of gslbservice pool groups and pool members from the gslb follower members. Field introduced in 20.1.5. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `error_resync_interval` - Frequency with which errored messages are resynced to follower sites. Value of 0 disables resync behavior. Allowed values are 60-3600. Special values are 0 - disable. Field introduced in 18.2.3. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `is_federated` - This field indicates that this object is replicated across gslb federation. Field introduced in 17.1.3. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `leader_cluster_uuid` - Mark this site as leader of gslb configuration. This site is the one among the avi sites. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `maintenance_mode` - This field disables the configuration operations on the leader for all federated objects. Cud operations on gslb, gslbservice, gslbgeodbprofile and other federated objects will be rejected. The rest-api disabling helps in upgrade scenarios where we don't want configuration sync operations to the gslb member when the member is being upgraded. This configuration programmatically blocks the leader from accepting new gslb configuration when member sites are undergoing upgrade. Field introduced in 17.2.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `name` - Name for the gslb object. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `replication_policy` - Policy for replicating configuration to the active follower sites. Field introduced in 20.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `send_interval` - Frequency with which group members communicate. Allowed values are 1-3600. Unit is sec. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `send_interval_prior_to_maintenance_mode` - The user can specify a send-interval while entering maintenance mode. The validity of this 'maintenance send-interval' is only during maintenance mode. When the user leaves maintenance mode, the original send-interval is reinstated. This internal variable is used to store the original send-interval. Field introduced in 18.2.3. Unit is sec. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `sites` - Select avi site member belonging to this gslb. Minimum of 1 items required. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_scoped` - This field indicates tenant visibility for gs pool member selection across the gslb federated objects.tenant scope can be set only during the gslb create and cannot be changed once it is set. Field introduced in 18.2.12,20.1.4. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `third_party_sites` - Third party site member belonging to this gslb. Field introduced in 17.1.1. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the gslb object. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `view_id` - The view-id is used in change-leader mode to differentiate partitioned groups while they have the same gslb namespace. Each partitioned group will be able to operate independently by using the view-id. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

