
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
page_title: "AVI: avi_cloudproperties"
sidebar_current: "docs-avi-datasource-cloudproperties"
description: |-
  Get information of Avi CloudProperties.
---

# avi_cloudproperties

This data source is used to to get avi_cloudproperties objects.

## Example Usage

```hcl
data "CloudProperties" "foo_CloudProperties" {
    uuid = "CloudProperties-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search CloudProperties by name.
* `uuid` - (Optional) Search CloudProperties by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cc_props` - Cloudconnector properties.
* `cc_vtypes` - Cloud types supported by cloudconnector.
* `hyp_props` - Hypervisor properties.
* `info` - Properties specific to a cloud type.
* `uuid` - General description.

