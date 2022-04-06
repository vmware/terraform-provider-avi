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

* `configpb_attributes` - Protobuf versioning for config pbs. Field introduced in 21.1.1. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `description` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `image_id` - Unique id of the amazon machine image (ami)  or openstack vm id. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `markers` - List of labels to be used for granular rbac. Field introduced in 20.1.5. Allowed in enterprise with any value edition, essentials with any value edition, basic with any value edition, enterprise with cloud services edition.
* `mesos` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `name` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `openstack` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `use_external_asg` - If set to true, serverautoscalepolicy will use the autoscaling group (external_autoscaling_groups) from pool to perform scale up and scale down. Pool should have single autoscaling group configured. Field introduced in 17.2.3. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uuid` - Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

