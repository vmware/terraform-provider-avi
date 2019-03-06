
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
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
data "PoolGroup" "foo_PoolGroup" {
    uuid = "PoolGroup-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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
* `fail_action` - Enable an action - close connection, http redirect, or local http response - when a pool group failure happens.
* `implicit_priority_labels` - Whether an implicit set of priority labels is generated.
* `members` - List of pool group members object of type poolgroupmember.
* `min_servers` - The minimum number of servers to distribute traffic to.
* `name` - The name of the pool group.
* `priority_labels_ref` - Uuid of the priority labels.
* `service_metadata` - Metadata pertaining to the service provided by this poolgroup.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the pool group.

