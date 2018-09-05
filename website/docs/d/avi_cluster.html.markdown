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
data "Cluster" "foo_Cluster" {
    uuid = "Cluster-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Cluster by name.
* `uuid` - (Optional) Search Cluster by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `name` - General description.
* `nodes` - General description.
* `rejoin_nodes_automatically` - Re-join cluster nodes automatically in the event one of the node is reset to factory.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.
* `virtual_ip` - A virtual ip address.

