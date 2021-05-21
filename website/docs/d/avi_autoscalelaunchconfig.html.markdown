<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_autoscalelaunchconfig"
sidebar_current: "docs-avi-datasource-autoscalelaunchconfig"
description: |-
  Get information of Avi AutoScaleLaunchConfig.
---

# avi_autoscalelaunchconfig

This data source is used to to get avi_autoscalelaunchconfig objects.

## Example Usage

```hcl
data "avi_autoscalelaunchconfig" "foo_autoscalelaunchconfig" {
    uuid = "autoscalelaunchconfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AutoScaleLaunchConfig by name.
* `uuid` - (Optional) Search AutoScaleLaunchConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1.
* `description` - User defined description for the object.
* `image_id` - Unique id of the amazon machine image (ami)  or openstack vm id.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5.
* `mesos` - Dict settings for autoscalelaunchconfig.
* `name` - Name of the object.
* `openstack` - Dict settings for autoscalelaunchconfig.
* `tenant_ref` - It is a reference to an object of type tenant.
* `use_external_asg` - If set to true, serverautoscalepolicy will use the autoscaling group (external_autoscaling_groups) from pool to perform scale up and scale down. Pool should have single autoscaling group configured. Field introduced in 17.2.3.
* `uuid` - Unique object identifier of the object.

