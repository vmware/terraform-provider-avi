############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

---
layout: "avi"
page_title: "AVI: avi_poolgroup"
sidebar_current: "docs-avi-datasource-poolgroup"
description: |-
  Get information of Avi PoolGroup.
---

# avi_poolgroup

This data source is used to to get avi_poolgroup objects.

## Example Usage

```hcl
data "avi_poolgroup" "foo_poolgroup" {
    uuid = "poolgroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search PoolGroup by name.
* `uuid` - (Optional) Search PoolGroup by uuid.
* `cloud_ref` - (Optional) Search PoolGroup by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_config_cksum` - Checksum of cloud configuration for poolgroup.
* `cloud_ref` - It is a reference to an object of type cloud.
* `created_by` - Name of the user who created the object.
* `deployment_policy_ref` - When setup autoscale manager will automatically promote new pools into production when deployment goals are met.
* `description` - Description of pool group.
* `enable_http2` - Enable http/2 for traffic from virtualservice to all the backend servers in all the pools configured under this poolgroup.
* `fail_action` - Enable an action - close connection, http redirect, or local http response - when a pool group failure happens.
* `implicit_priority_labels` - Whether an implicit set of priority labels is generated.
* `labels` - Key value pairs for granular object access control.
* `members` - List of pool group members object of type poolgroupmember.
* `min_servers` - The minimum number of servers to distribute traffic to.
* `name` - The name of the pool group.
* `priority_labels_ref` - Uuid of the priority labels.
* `service_metadata` - Metadata pertaining to the service provided by this poolgroup.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the pool group.

