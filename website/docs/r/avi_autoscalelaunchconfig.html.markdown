<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_autoscalelaunchconfig"
sidebar_current: "docs-avi-resource-autoscalelaunchconfig"
description: |-
  Creates and manages Avi AutoScaleLaunchConfig.
---

# avi_autoscalelaunchconfig

The AutoScaleLaunchConfig resource allows the creation and management of Avi AutoScaleLaunchConfig

## Example Usage

```hcl
resource "avi_autoscalelaunchconfig" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the object.
* `configpb_attributes` - (Optional) Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - (Optional) User defined description for the object.
* `image_id` - (Optional) Unique id of the amazon machine image (ami)  or openstack vm id.
* `markers` - (Optional) List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `mesos` - (Optional) Dict settings for autoscalelaunchconfig.
* `openstack` - (Optional) Dict settings for autoscalelaunchconfig.
* `tenant_ref` - (Optional) It is a reference to an object of type tenant.
* `use_external_asg` - (Optional) If set to true, serverautoscalepolicy will use the autoscaling group (external_autoscaling_groups) from pool to perform scale up and scale down. Pool should have single autoscaling group configured. Field introduced in 17.2.3.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Unique object identifier of the object.

