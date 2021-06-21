<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_serverautoscalepolicy"
sidebar_current: "docs-avi-resource-serverautoscalepolicy"
description: |-
  Creates and manages Avi ServerAutoScalePolicy.
---

# avi_serverautoscalepolicy

The ServerAutoScalePolicy resource allows the creation and management of Avi ServerAutoScalePolicy

## Example Usage

```hcl
resource "avi_serverautoscalepolicy" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `delay_for_server_garbage_collection` - (Optional) Delay in minutes after which a down server will be removed from pool. Value 0 disables this functionality. Field introduced in 20.1.3.
* `description` - (Optional) User defined description for the object.
* `intelligent_autoscale` - (Optional) Use avi intelligent autoscale algorithm where autoscale is performed by comparing load on the pool against estimated capacity of all the servers.
* `intelligent_scalein_margin` - (Optional) Maximum extra capacity as percentage of load used by the intelligent scheme. Scale-in is triggered when available capacity is more than this margin. Allowed values are 1-99.
* `intelligent_scaleout_margin` - (Optional) Minimum extra capacity as percentage of load used by the intelligent scheme. Scale-out is triggered when available capacity is less than this margin. Allowed values are 1-99.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in basic edition, essentials edition, enterprise edition.
* `max_scalein_adjustment_step` - (Optional) Maximum number of servers to scale-in simultaneously. The actual number of servers to scale-in is chosen such that target number of servers is always more than or equal to the min_size.
* `max_scaleout_adjustment_step` - (Optional) Maximum number of servers to scale-out simultaneously. The actual number of servers to scale-out is chosen such that target number of servers is always less than or equal to the max_size.
* `max_size` - (Optional) Maximum number of servers after scale-out. Allowed values are 0-400.
* `min_size` - (Optional) No scale-in happens once number of operationally up servers reach min_servers. Allowed values are 0-400.
* `scalein_alertconfig_refs` - (Optional) Trigger scale-in when alerts due to any of these alert configurations are raised. It is a reference to an object of type alertconfig.
* `scalein_cooldown` - (Optional) Cooldown period during which no new scale-in is triggered to allow previous scale-in to successfully complete. Unit is sec.
* `scaleout_alertconfig_refs` - (Optional) Trigger scale-out when alerts due to any of these alert configurations are raised. It is a reference to an object of type alertconfig.
* `scaleout_cooldown` - (Optional) Cooldown period during which no new scale-out is triggered to allow previous scale-out to successfully complete. Unit is sec.
* `scheduled_scalings` - (Optional) Schedule-based scale-in/out policy. During schedule intervals, metrics based autoscale is not enabled and number of servers will be solely derived from schedulescale policy. Field introduced in 21.1.1. Maximum of 1 items allowed.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `use_predicted_load` - (Optional) Use predicted load rather than current load.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

