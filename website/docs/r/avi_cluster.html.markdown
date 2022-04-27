<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_cluster"
sidebar_current: "docs-avi-resource-cluster"
description: |-
  Creates and manages Avi Cluster.
---

# avi_cluster

The Cluster resource allows the creation and management of Avi Cluster

## Example Usage

```hcl
resource "avi_cluster" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `nodes` - (Required) Minimum of 1 items required. Maximum of 7 items allowed. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `rejoin_nodes_automatically` - (Optional) Re-join cluster nodes automatically in the event one of the node is reset to factory. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `virtual_ip` - (Optional) A virtual ip address. This ip address will be dynamically reconfigured so that it always is the ip of the cluster leader. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

