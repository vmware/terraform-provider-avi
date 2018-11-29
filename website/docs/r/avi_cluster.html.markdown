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
resource "Cluster" "foo" {
    name = "terraform-example-foo"
    tenant = "admin"
}
```

## Argument Reference

The following arguments are supported:

    * `name` - (Required) argument_description.
        * `nodes` - (Optional ) argument_description.
        * `rejoin_nodes_automatically` - (Optional ) argument_description.
        * `tenant_ref` - (Optional ) argument_description.
            * `virtual_ip` - (Optional ) argument_description.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

                    * `uuid` - argument_description.
