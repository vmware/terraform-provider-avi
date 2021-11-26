############################################################################
# ========================================================================
# Copyright 2021 VMware, Inc.  All rights reserved. VMware Confidential
# ========================================================================
###

<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_image"
sidebar_current: "docs-avi-datasource-image"
description: |-
  Get information of Avi Image.
---

# avi_image

This data source is used to to get avi_image objects.

## Example Usage

```hcl
data "avi_image" "foo_image" {
    uuid = "image-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search Image by name.
* `uuid` - (Optional) Search Image by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `cloud_info_values` - This field describes the cloud info specific to the base image. Field introduced in 20.1.1.
* `controller_info` - Controller package details. Field introduced in 18.2.6.
* `controller_patch_name` - Mandatory controller patch name that is applied along with this base image. Field introduced in 18.2.10, 20.1.1.
* `controller_patch_uuid` - It references the controller-patch associated with the uber image. Field introduced in 18.2.8, 20.1.1.
* `duration` - Time taken to upload the image in seconds. Field introduced in 21.1.3. Unit is sec.
* `end_time` - Image upload end time. Field introduced in 21.1.3.
* `events` - Image events for image upload operation. Field introduced in 21.1.3.
* `img_state` - Status of the image. Field introduced in 21.1.3.
* `migrations` - This field describes the api migration related information. Field introduced in 18.2.6.
* `name` - Name of the image. Field introduced in 18.2.6.
* `progress` - Image upload progress which holds value between 0-100. Allowed values are 0-100. Field introduced in 21.1.3. Unit is percent.
* `se_info` - Se package details. Field introduced in 18.2.6.
* `se_patch_name` - Mandatory serviceengine patch name that is applied along with this base image. Field introduced in 18.2.10, 20.1.1.
* `se_patch_uuid` - It references the service engine patch associated with the uber image. Field introduced in 18.2.8, 20.1.1.
* `start_time` - Image upload start time. Field introduced in 21.1.3.
* `tasks_completed` - Completed set of tasks for image upload. Field introduced in 21.1.3.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.6.
* `total_tasks` - Total number of tasks for image upload. Field introduced in 21.1.3.
* `type` - Type of the image patch/system. Enum options - IMAGE_TYPE_PATCH, IMAGE_TYPE_SYSTEM, IMAGE_TYPE_MUST_CHECK. Field introduced in 18.2.6.
* `uber_bundle` - Status to check if the image is an uber bundle. Field introduced in 18.2.8, 20.1.1.
* `uuid` - Uuid of the image. Field introduced in 18.2.6.

