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
page_title: "AVI: avi_cluster"
sidebar_current: "docs-avi-datasource-cluster"
description: |-
  Get information of Avi Cluster.
---

# avi_cluster

This data source is used to to get avi_cluster objects.

## Example Usage

```hcl
data "avi_cluster" "foo_cluster" {
    uuid = "cluster-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Cluster by name.
* `uuid` - (Optional) Search Cluster by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - Name of the object.
* `nodes` - Minimum of 1 items required. Maximum of 7 items allowed.
* `rejoin_nodes_automatically` - Re-join cluster nodes automatically in the event one of the node is reset to factory.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.
* `virtual_ip` - A virtual ip address. This ip address will be dynamically reconfigured so that it always is the ip of the cluster leader.

