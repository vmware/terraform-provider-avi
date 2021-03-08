############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `availability_zone` - Placeholder for description of property availability_zone of obj type serviceengine field type string  type str.
* `cloud_ref` - It is a reference to an object of type cloud.
* `container_mode` - Boolean flag to set container_mode.
* `container_type` - Enum options - container_type_bridge, container_type_host, container_type_host_dpdk.
* `controller_created` - Boolean flag to set controller_created.
* `controller_ip` - Placeholder for description of property controller_ip of obj type serviceengine field type string  type str.
* `data_vnics` - List of list.
* `enable_state` - Inorder to disable se set this field appropriately. Enum options - SE_STATE_ENABLED, SE_STATE_DISABLED_FOR_PLACEMENT, SE_STATE_DISABLED, SE_STATE_DISABLED_FORCE.
* `flavor` - Placeholder for description of property flavor of obj type serviceengine field type string  type str.
* `host_ref` - It is a reference to an object of type vimgrhostruntime.
* `hypervisor` - Enum options - default, vmware_esx, kvm, vmware_vsan, xen.
* `mgmt_vnic` - Dict settings for serviceengine.
* `name` - Name of the object.
* `resources` - Dict settings for serviceengine.
* `se_group_ref` - It is a reference to an object of type serviceenginegroup.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - Unique object identifier of the object.

