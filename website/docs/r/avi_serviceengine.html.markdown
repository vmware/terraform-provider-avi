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

* `availability_zone` - (Optional) Placeholder for description of property availability_zone of obj type serviceengine field type string  type str.
* `cloud_ref` - (Optional) It is a reference to an object of type cloud.
* `container_mode` - (Optional) Boolean flag to set container_mode.
* `container_type` - (Optional) Enum options - container_type_bridge, container_type_host, container_type_host_dpdk.
* `controller_created` - (Optional) Boolean flag to set controller_created.
* `controller_ip` - (Optional) Placeholder for description of property controller_ip of obj type serviceengine field type string  type str.
* `data_vnics` - (Optional) List of list.
* `enable_state` - (Optional) Inorder to disable se set this field appropriately.
* `flavor` - (Optional) Placeholder for description of property flavor of obj type serviceengine field type string  type str.
* `host_ref` - (Optional) It is a reference to an object of type vimgrhostruntime.
* `hypervisor` - (Optional) Enum options - default, vmware_esx, kvm, vmware_vsan, xen.
* `mgmt_vnic` - (Optional) Dict settings for serviceengine.
* `name` - (Optional) Name of the object.
* `resources` - (Optional) Dict settings for serviceengine.
* `se_group_ref` - (Optional) It is a reference to an object of type serviceenginegroup.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

