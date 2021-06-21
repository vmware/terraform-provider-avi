<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_prioritylabels"
sidebar_current: "docs-avi-resource-prioritylabels"
description: |-
  Creates and manages Avi PriorityLabels.
---

# avi_prioritylabels

The PriorityLabels resource allows the creation and management of Avi PriorityLabels

## Example Usage

```hcl
resource "avi_prioritylabels" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the priority labels.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - (Optional) A description of the priority labels.
* `equivalent_labels` - (Optional) Equivalent priority labels in descending order.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the priority labels.

