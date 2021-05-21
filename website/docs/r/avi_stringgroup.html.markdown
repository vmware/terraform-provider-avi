<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_stringgroup"
sidebar_current: "docs-avi-resource-stringgroup"
description: |-
  Creates and manages Avi StringGroup.
---

# avi_stringgroup

The StringGroup resource allows the creation and management of Avi StringGroup

## Example Usage

```hcl
resource "avi_stringgroup" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the string group.
* `type` - (Required) Type of stringgroup. Enum options - SG_TYPE_STRING, SG_TYPE_KEYVAL.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - (Optional) User defined description for the object.
* `kv` - (Optional) Configure key value in the string group.
* `longest_match` - (Optional) Enable the longest match, default is the shortest match. Field introduced in 18.2.8.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the string group.

