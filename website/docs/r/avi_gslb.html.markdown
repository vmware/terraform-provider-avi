############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_gslb"
sidebar_current: "docs-avi-resource-gslb"
description: |-
  Creates and manages Avi Gslb.
---

# avi_gslb

The Gslb resource allows the creation and management of Avi Gslb

## Example Usage

```hcl
resource "avi_gslb" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name for the gslb object.
* `async_interval` - (Optional) Frequency with which messages are propagated to vs mgr.
* `clear_on_max_retries` - (Optional) Max retries after which the remote site is treated as a fresh start.
* `client_ip_addr_group` - (Optional) Group to specify if the client ip addresses are public or private.
* `description` - (Optional) User defined description for the object.
* `dns_configs` - (Optional) Sub domain configuration for the gslb.
* `error_resync_interval` - (Optional) Frequency with which errored messages are resynced to follower sites.
* `is_federated` - (Optional) This field indicates that this object is replicated across gslb federation.
* `leader_cluster_uuid` - (Optional) Mark this site as leader of gslb configuration.
* `maintenance_mode` - (Optional) This field disables the configuration operations on the leader for all federated objects.
* `send_interval` - (Optional) Frequency with which group members communicate.
* `send_interval_prior_to_maintenance_mode` - (Optional) The user can specify a send-interval while entering maintenance mode.
* `sites` - (Optional) Select avi site member belonging to this gslb.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `third_party_sites` - (Optional) Third party site member belonging to this gslb.
* `view_id` - (Optional) The view-id is used in change-leader mode to differentiate partitioned groups while they have the same gslb namespace.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the gslb object.

