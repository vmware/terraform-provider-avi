
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
page_title: "AVI: avi_vrfcontext"
sidebar_current: "docs-avi-datasource-vrfcontext"
description: |-
  Get information of Avi VrfContext.
---

# avi_vrfcontext

This data source is used to to get avi_vrfcontext objects.

## Example Usage

```hcl
data "VrfContext" "foo_VrfContext" {
    uuid = "VrfContext-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search VrfContext by name.
* `uuid` - (Optional) Search VrfContext by uuid.
* `cloud_ref` - (Optional) Search VrfContext by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `bgp_profile` - Bgp local and peer info.
* `cloud_ref` - It is a reference to an object of type cloud.
* `debugvrfcontext` - Configure debug flags for vrf.
* `description` - General description.
* `gateway_mon` - Configure ping based heartbeat check for gateway in service engines of vrf.
* `internal_gateway_monitor` - Configure ping based heartbeat check for all default gateways in service engines of vrf.
* `name` - General description.
* `static_routes` - General description.
* `system_default` - General description.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.

