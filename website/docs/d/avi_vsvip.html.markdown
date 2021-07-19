<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_vsvip"
sidebar_current: "docs-avi-datasource-vsvip"
description: |-
  Get information of Avi VsVip.
---

# avi_vsvip

This data source is used to to get avi_vsvip objects.

## Example Usage

```hcl
data "avi_vsvip" "foo_vsvip" {
    uuid = "vsvip-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search VsVip by name.
* `uuid` - (Optional) Search VsVip by uuid.
* `cloud_ref` - (Optional) Search VsVip by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_ref` - It is a reference to an object of type cloud. Field introduced in 17.1.1.
* `dns_info` - Service discovery specific data including fully qualified domain name, type and time-to-live of the dns record. Field introduced in 17.1.1.
* `east_west_placement` - Force placement on all service engines in the service engine group (container clouds only). Field introduced in 17.1.1.
* `name` - Name for the vsvip object. Field introduced in 17.1.1.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 17.1.1.
* `use_standard_alb` - This overrides the cloud level default and needs to match the se group value in which it will be used if the se group use_standard_alb value is set. This is only used when fip is used for vs on azure cloud. Field introduced in 18.2.3.
* `uuid` - Uuid of the vsvip object. Field introduced in 17.1.1.
* `vip` - List of virtual service ips and other shareable entities. Field introduced in 17.1.1.
* `vrf_context_ref` - Virtual routing context that the virtual service is bound to. This is used to provide the isolation of the set of networks the application is attached to. It is a reference to an object of type vrfcontext. Field introduced in 17.1.1.
* `vsvip_cloud_config_cksum` - Checksum of cloud configuration for vsvip. Internally set by cloud connector. Field introduced in 17.2.9, 18.1.2.

