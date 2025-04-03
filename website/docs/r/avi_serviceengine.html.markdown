<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_serviceengine"
sidebar_current: "docs-avi-resource-serviceengine"
description: |-
  Creates and manages Avi ServiceEngine.
---

# avi_serviceengine

The ServiceEngine resource allows the creation and management of Avi ServiceEngine

## Example Usage

```hcl
resource "avi_serviceengine" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `availability_zone` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `container_mode` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `container_type` - (Optional) Enum options - container_type_bridge, container_type_host, container_type_host_dpdk. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `controller_created` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `controller_ip` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `data_vnics` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `enable_state` - (Optional) Inorder to disable se set this field appropriately. Enum options - SE_STATE_ENABLED, SE_STATE_DISABLED_FOR_PLACEMENT, SE_STATE_DISABLED, SE_STATE_DISABLED_FORCE, SE_STATE_DISABLED_WITH_SCALEIN, SE_STATE_DISABLED_NO_TRAFFIC, SE_STATE_DISABLED_FORCE_WITH_MIGRATE. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `flavor` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `host_ref` - (Optional) It is a reference to an object of type vimgrhostruntime. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `hypervisor` - (Optional) Enum options - default, vmware_esx, kvm, vmware_vsan, xen. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `mgmt_vnic` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `nsxt_no_hotplug` - (Optional) If set to true, controller does not hotplugg the vnics. Field introduced in 30.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `resources` - (Optional) Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `se_group_ref` - (Optional) It is a reference to an object of type serviceenginegroup. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

