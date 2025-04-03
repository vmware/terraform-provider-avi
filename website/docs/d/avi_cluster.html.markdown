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

* `name` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `nodes` - Minimum of 1 items required. Maximum of 7 items allowed. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `rejoin_nodes_automatically` - Re-join cluster nodes automatically in the event one of the node is reset to factory. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `virtual_ip` - A v4 virtual ip address for the cluster that always points to the v4 ip of the leader node in cluster. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `virtual_ip6` - A v6 virtual ip address for the cluster that always points to the v6 ip of the leader node in cluster. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

