
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
page_title: "AVI: avi_gslbgeodbprofile"
sidebar_current: "docs-avi-datasource-gslbgeodbprofile"
description: |-
  Get information of Avi GslbGeoDbProfile.
---

# avi_gslbgeodbprofile

This data source is used to to get avi_gslbgeodbprofile objects.

## Example Usage

```hcl
data "GslbGeoDbProfile" "foo_GslbGeoDbProfile" {
    uuid = "GslbGeoDbProfile-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search GslbGeoDbProfile by name.
* `uuid` - (Optional) Search GslbGeoDbProfile by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - Field introduced in 17.1.1.
* `entries` - List of geodb entries.
* `is_federated` - This field indicates that this object is replicated across gslb federation.
* `name` - A user-friendly name for the geodb profile.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Uuid of the geodb profile.

