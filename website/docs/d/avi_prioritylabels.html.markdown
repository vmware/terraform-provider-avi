
############################################################################
#
# AVI CONFIDENTIAL
# __________________
#
# [2013] - [2019] Avi Networks Incorporated
# All Rights Reserved.
#
# NOTICE: All information contained herein is, and remains the property
# of Avi Networks Incorporated and its suppliers, if any. The intellectual
# and technical concepts contained herein are proprietary to Avi Networks
# Incorporated, and its suppliers and are covered by U.S. and Foreign
# Patents, patents in process, and are protected by trade secret or
# copyright law, and other laws. Dissemination of this information or
# reproduction of this material is strictly forbidden unless prior written
# permission is obtained from Avi Networks Incorporated.
###

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
data "PriorityLabels" "foo_PriorityLabels" {
    uuid = "PriorityLabels-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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

* `cloud_ref` - It is a reference to an object of type cloud.
* `description` - A description of the priority labels.
* `equivalent_labels` - Equivalent priority labels in descending order.
* `name` - The name of the priority labels.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the priority labels.

