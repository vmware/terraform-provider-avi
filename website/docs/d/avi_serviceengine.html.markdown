
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
page_title: "AVI: avi_serviceengine"
sidebar_current: "docs-avi-datasource-serviceengine"
description: |-
  Get information of Avi ServiceEngine.
---

# avi_serviceengine

This data source is used to to get avi_serviceengine objects.

## Example Usage

```hcl
data "ServiceEngine" "foo_ServiceEngine" {
    uuid = "ServiceEngine-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
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

* `availability_zone` - General description.
* `cloud_ref` - It is a reference to an object of type cloud.
* `container_mode` - General description.
* `container_type` - Enum options - container_type_bridge, container_type_host, container_type_host_dpdk.
* `controller_created` - General description.
* `controller_ip` - General description.
* `data_vnics` - General description.
* `enable_state` - Inorder to disable se set this field appropriately.
* `flavor` - General description.
* `host_ref` - It is a reference to an object of type vimgrhostruntime.
* `hypervisor` - Enum options - default, vmware_esx, kvm, vmware_vsan, xen.
* `mgmt_vnic` - General description.
* `name` - General description.
* `resources` - General description.
* `se_group_ref` - It is a reference to an object of type serviceenginegroup.
* `tenant_ref` - It is a reference to an object of type tenant.
* `uuid` - General description.

