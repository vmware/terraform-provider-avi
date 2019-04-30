---
layout: "avi"
page_title: "AVI: avi_serverautoscalepolicy"
sidebar_current: "docs-avi-datasource-serverautoscalepolicy"
description: |-
  Get information of Avi ServerAutoScalePolicy.
---

# avi_serverautoscalepolicy

This data source is used to to get avi_serverautoscalepolicy objects.

## Example Usage

```hcl
data "ServerAutoScalePolicy" "foo_ServerAutoScalePolicy" {
    uuid = "ServerAutoScalePolicy-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ServerAutoScalePolicy by name.
* `uuid` - (Optional) Search ServerAutoScalePolicy by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - General description.
* `intelligent_autoscale` - Use avi intelligent autoscale algorithm where autoscale is performed by comparing load on the pool against estimated capacity of all the servers.
* `intelligent_scalein_margin` - Maximum extra capacity as percentage of load used by the intelligent scheme.
* `intelligent_scaleout_margin` - Minimum extra capacity as percentage of load used by the intelligent scheme.
* `max_scalein_adjustment_step` - Maximum number of servers to scalein simultaneously.
* `max_scaleout_adjustment_step` - Maximum number of servers to scaleout simultaneously.
* `max_size` - Maximum number of servers after scaleout.
* `min_size` - No scale-in happens once number of operationally up servers reach min_servers.
* `name` - General description.
* `scalein_alertconfig_refs` - Trigger scalein when alerts due to any of these alert configurations are raised.
* `scalein_cooldown` - Cooldown period during which no new scalein is triggered to allow previous scalein to successfully complete.
* `scaleout_alertconfig_refs` - Trigger scaleout when alerts due to any of these alert configurations are raised.
* `scaleout_cooldown` - Cooldown period during which no new scaleout is triggered to allow previous scaleout to successfully complete.
* `tenant_ref` - It is a reference to an object of type tenant.
* `use_predicted_load` - Use predicted load rather than current load.
* `uuid` - General description.

