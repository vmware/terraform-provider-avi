############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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
* `cloud_config_cksum` - (Optional) Checksum of cloud configuration for poolgroup. Internally set by cloud connector.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `created_by` - (Optional) Name of the user who created the object.
* `deactivate_primary_pool_on_down` - (Optional) Deactivate primary pool for selection when down until it is activated by user via clear poolgroup command. Field introduced in 20.1.7, 21.1.2, 21.1.3.
* `deployment_policy_ref` - (Optional) When setup autoscale manager will automatically promote new pools into production when deployment goals are met. It is a reference to an object of type poolgroupdeploymentpolicy.
* `description` - (Optional) Description of pool group.
* `enable_http2` - (Optional) Enable http/2 for traffic from virtualservice to all the backend servers in all the pools configured under this poolgroup. Field introduced in 20.1.1. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `fail_action` - (Optional) Enable an action - close connection, http redirect, or local http response - when a pool group failure happens. By default, a connection will be closed, in case the pool group experiences a failure.
* `implicit_priority_labels` - (Optional) Whether an implicit set of priority labels is generated. Field introduced in 17.1.9,17.2.3.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `members` - (Optional) List of pool group members object of type poolgroupmember.
* `min_servers` - (Optional) The minimum number of servers to distribute traffic to. Allowed values are 1-65535. Special values are 0 - 'disable'. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition.
* `priority_labels_ref` - (Optional) Uuid of the priority labels. If not provided, pool group member priority label will be interpreted as a number with a larger number considered higher priority. It is a reference to an object of type prioritylabels.
* `service_metadata` - (Optional) Metadata pertaining to the service provided by this poolgroup. In openshift/kubernetes environments, app metadata info is stored. Any user input to this field will be overwritten by avi vantage. Field introduced in 17.2.14,18.1.5,18.2.1.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the pool group.

