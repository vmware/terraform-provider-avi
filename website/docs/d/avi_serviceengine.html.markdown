<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_serviceengine"
sidebar_current: "docs-avi-datasource-serviceengine"
description: |-
  Get information of Avi ServiceEngine.
---

# avi_serviceengine

This data source is used to to get avi_serviceengine objects.

## Example Usage

```hcl
data "avi_serviceengine" "foo_serviceengine" {
    uuid = "serviceengine-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
    cloud_ref = "/api/cloud/?tenant=admin&name=Default-Cloud"
  }
```

## Argument Reference

* `name` - (Optional) Search ServiceEngine by name.
* `uuid` - (Optional) Search ServiceEngine by uuid.
* `cloud_ref` - (Optional) Search ServiceEngine by cloud_ref.
  
## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `availability_zone` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `cloud_ref` - It is a reference to an object of type cloud. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `container_mode` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `container_type` - Enum options - container_type_bridge, container_type_host, container_type_host_dpdk. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `controller_created` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `controller_ip` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `data_vnics` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `enable_state` - Inorder to disable se set this field appropriately. Enum options - SE_STATE_ENABLED, SE_STATE_DISABLED_FOR_PLACEMENT, SE_STATE_DISABLED, SE_STATE_DISABLED_FORCE. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `flavor` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `host_ref` - It is a reference to an object of type vimgrhostruntime. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `hypervisor` - Enum options - default, vmware_esx, kvm, vmware_vsan, xen. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `mgmt_vnic` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `name` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `resources` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `se_group_ref` - It is a reference to an object of type serviceenginegroup. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

