############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "Avi: avi_poolgroup"
sidebar_current: "docs-avi-resource-poolgroup"
description: |-
  Creates and manages Avi PoolGroup.
---

# avi_poolgroup

The PoolGroup resource allows the creation and management of Avi PoolGroup

## Example Usage

```hcl
resource "avi_poolgroup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the pool group.
* `cloud_config_cksum` - (Optional) Checksum of cloud configuration for poolgroup.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `created_by` - (Optional) Name of the user who created the object.
* `deployment_policy_ref` - (Optional) When setup autoscale manager will automatically promote new pools into production when deployment goals are met.
* `description` - (Optional) Description of pool group.
* `fail_action` - (Optional) Enable an action - close connection, http redirect, or local http response - when a pool group failure happens.
* `implicit_priority_labels` - (Optional) Whether an implicit set of priority labels is generated.
* `labels` - (Optional) Key value pairs for granular object access control.
* `members` - (Optional) List of pool group members object of type poolgroupmember.
* `min_servers` - (Optional) The minimum number of servers to distribute traffic to.
* `priority_labels_ref` - (Optional) Uuid of the priority labels.
* `service_metadata` - (Optional) Metadata pertaining to the service provided by this poolgroup.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the pool group.

