<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_prioritylabels"
sidebar_current: "docs-avi-datasource-prioritylabels"
description: |-
  Get information of Avi PriorityLabels.
---

# avi_prioritylabels

This data source is used to to get avi_prioritylabels objects.

## Example Usage

```hcl
data "avi_prioritylabels" "foo_prioritylabels" {
    uuid = "prioritylabels-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search PriorityLabels by name.
* `uuid` - (Optional) Search PriorityLabels by uuid.
* `cloud_ref` - (Optional) Search PriorityLabels by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_ref` - It is a reference to an object of type cloud. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `description` - A description of the priority labels. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `equivalent_labels` - Equivalent priority labels in descending order. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise edition with any value, essentials edition with any value, basic edition with any value, enterprise with cloud services edition.
* `name` - The name of the priority labels. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the priority labels. Allowed in enterprise edition with any value, essentials, basic, enterprise with cloud services edition.

