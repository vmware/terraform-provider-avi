<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_geodb"
sidebar_current: "docs-avi-resource-geodb"
description: |-
  Creates and manages Avi GeoDB.
---

# avi_geodb

The GeoDB resource allows the creation and management of Avi GeoDB

## Example Usage

```hcl
resource "avi_geodb" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `files` - (Required) Geo database files. Field introduced in 21.1.1.
* `name` - (Required) Geo database name. Field introduced in 21.1.1.
* `description` - (Optional) Description. Field introduced in 21.1.1.
* `is_federated` - (Optional) This field indicates that this object is replicated across gslb federation. Field introduced in 21.1.1.
* `mappings` - (Optional) Custom mappings of geo values. All mappings which start with the prefix 'system-' (any case) are reserved for system default objects and may be overwritten. Field introduced in 21.1.1.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 21.1.1.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of this object. Field introduced in 21.1.1.

