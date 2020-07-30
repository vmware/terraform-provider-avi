############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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
* `description` - (Optional) User defined description for the object.
* `intelligent_autoscale` - (Optional) Use avi intelligent autoscale algorithm where autoscale is performed by comparing load on the pool against estimated capacity of all the servers.
* `intelligent_scalein_margin` - (Optional) Maximum extra capacity as percentage of load used by the intelligent scheme.
* `intelligent_scaleout_margin` - (Optional) Minimum extra capacity as percentage of load used by the intelligent scheme.
* `max_scalein_adjustment_step` - (Optional) Maximum number of servers to scalein simultaneously.
* `max_scaleout_adjustment_step` - (Optional) Maximum number of servers to scaleout simultaneously.
* `max_size` - (Optional) Maximum number of servers after scaleout.
* `min_size` - (Optional) No scale-in happens once number of operationally up servers reach min_servers.
* `scalein_alertconfig_refs` - (Optional) Trigger scalein when alerts due to any of these alert configurations are raised.
* `scalein_cooldown` - (Optional) Cooldown period during which no new scalein is triggered to allow previous scalein to successfully complete.
* `scaleout_alertconfig_refs` - (Optional) Trigger scaleout when alerts due to any of these alert configurations are raised.
* `scaleout_cooldown` - (Optional) Cooldown period during which no new scaleout is triggered to allow previous scaleout to successfully complete.
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

