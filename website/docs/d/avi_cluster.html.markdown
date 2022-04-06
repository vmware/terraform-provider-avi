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

* `name` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `nodes` - Minimum of 1 items required. Maximum of 7 items allowed. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `rejoin_nodes_automatically` - Re-join cluster nodes automatically in the event one of the node is reset to factory. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `virtual_ip` - A virtual ip address. This ip address will be dynamically reconfigured so that it always is the ip of the cluster leader. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

