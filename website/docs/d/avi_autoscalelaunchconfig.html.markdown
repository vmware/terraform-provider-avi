
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
page_title: "AVI: avi_autoscalelaunchconfig"
sidebar_current: "docs-avi-datasource-autoscalelaunchconfig"
description: |-
  Get information of Avi AutoScaleLaunchConfig.
---

# avi_autoscalelaunchconfig

This data source is used to to get avi_autoscalelaunchconfig objects.

## Example Usage

```hcl
data "AutoScaleLaunchConfig" "foo_AutoScaleLaunchConfig" {
    uuid = "AutoScaleLaunchConfig-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search AutoScaleLaunchConfig by name.
* `uuid` - (Optional) Search AutoScaleLaunchConfig by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `description` - General description.
* `image_id` - Unique id of the amazon machine image (ami)  or openstack vm id.
* `mesos` - General description.
* `name` - General description.
* `openstack` - General description.
* `tenant_ref` - It is a reference to an object of type tenant.
* `use_external_asg` - If set to true, serverautoscalepolicy will use the autoscaling group (external_autoscaling_groups) from pool to perform scale up and scale down.
* `uuid` - General description.

