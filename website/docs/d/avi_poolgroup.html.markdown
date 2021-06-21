<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
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

* `cloud_config_cksum` - Checksum of cloud configuration for poolgroup. Internally set by cloud connector.
* `cloud_ref` - It is a reference to an object of type cloud.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `created_by` - Name of the user who created the object.
* `deployment_policy_ref` - When setup autoscale manager will automatically promote new pools into production when deployment goals are met. It is a reference to an object of type poolgroupdeploymentpolicy.
* `description` - Description of pool group.
* `enable_http2` - Enable http/2 for traffic from virtualservice to all the backend servers in all the pools configured under this poolgroup. Field introduced in 20.1.1. Allowed in basic(allowed values- false) edition, essentials(allowed values- false) edition, enterprise edition.
* `fail_action` - Enable an action - close connection, http redirect, or local http response - when a pool group failure happens. By default, a connection will be closed, in case the pool group experiences a failure.
* `implicit_priority_labels` - Whether an implicit set of priority labels is generated. Field introduced in 17.1.9,17.2.3.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `members` - List of pool group members object of type poolgroupmember.
* `min_servers` - The minimum number of servers to distribute traffic to. Allowed values are 1-65535. Special values are 0 - 'disable'. Allowed in basic(allowed values- 0) edition, essentials(allowed values- 0) edition, enterprise edition.
* `name` - The name of the pool group.
* `priority_labels_ref` - Uuid of the priority labels. If not provided, pool group member priority label will be interpreted as a number with a larger number considered higher priority. It is a reference to an object of type prioritylabels.
* `service_metadata` - Metadata pertaining to the service provided by this poolgroup. In openshift/kubernetes environments, app metadata info is stored. Any user input to this field will be overwritten by avi vantage. Field introduced in 17.2.14,18.1.5,18.2.1.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the pool group.

