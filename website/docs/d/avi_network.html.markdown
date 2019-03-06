
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
page_title: "AVI: avi_network"
sidebar_current: "docs-avi-datasource-network"
description: |-
  Get information of Avi Network.
---

# avi_network

This data source is used to to get avi_network objects.

## Example Usage

```hcl
data "Network" "foo_Network" {
    uuid = "Network-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search Network by name.
* `uuid` - (Optional) Search Network by uuid.
* `cloud_ref` - (Optional) Search Network by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_ref` - It is a reference to an object of type cloud.
* `configured_subnets` - General description.
* `dhcp_enabled` - Select the ip address management scheme for this network.
* `exclude_discovered_subnets` - When selected, excludes all discovered subnets in this network from consideration for virtual service placement.
* `ip6_autocfg_enabled` - Enable ipv6 auto configuration.
* `name` - General description.
* `synced_from_se` - General description.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.
* `vcenter_dvs` - General description.
* `vrf_context_ref` - It is a reference to an object of type vrfcontext.

