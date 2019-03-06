
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
page_title: "AVI: avi_networksecuritypolicy"
sidebar_current: "docs-avi-datasource-networksecuritypolicy"
description: |-
  Get information of Avi NetworkSecurityPolicy.
---

# avi_networksecuritypolicy

This data source is used to to get avi_networksecuritypolicy objects.

## Example Usage

```hcl
data "NetworkSecurityPolicy" "foo_NetworkSecurityPolicy" {
    uuid = "NetworkSecurityPolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search NetworkSecurityPolicy by name.
* `uuid` - (Optional) Search NetworkSecurityPolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_config_cksum` - Checksum of cloud configuration for network sec policy.
* `created_by` - Creator name.
* `description` - General description.
* `name` - General description.
* `rules` - General description.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.

