<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_image"
sidebar_current: "docs-avi-resource-image"
description: |-
  Creates and manages Avi Image.
---

# avi_image

The Image resource allows the creation and management of Avi Image

## Example Usage

```hcl
resource "avi_image" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of the image. Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `cloud_info_values` - (Optional) This field describes the cloud info specific to the base image. Field introduced in 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `controller_info` - (Optional) Controller package details. Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `controller_patch_name` - (Optional) Mandatory controller patch name that is applied along with this base image. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `controller_patch_ref` - (Optional) It references the controller-patch associated with the uber image. It is a reference to an object of type image. Field introduced in 18.2.8, 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `duration` - (Optional) Time taken to upload the image in seconds. Field introduced in 21.1.3. Unit is sec. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `end_time` - (Optional) Image upload end time. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `events` - (Optional) Image events for image upload operation. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `img_state` - (Optional) Status of the image. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `migrations` - (Optional) This field describes the api migration related information. Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `progress` - (Optional) Image upload progress which holds value between 0-100. Allowed values are 0-100. Field introduced in 21.1.3. Unit is percent. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `se_info` - (Optional) Se package details. Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `se_patch_name` - (Optional) Mandatory serviceengine patch name that is applied along with this base image. Field introduced in 18.2.10, 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `se_patch_ref` - (Optional) It references the service engine patch associated with the uber image. It is a reference to an object of type image. Field introduced in 18.2.8, 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `start_time` - (Optional) Image upload start time. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `tasks_completed` - (Optional) Completed set of tasks for image upload. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `total_tasks` - (Optional) Total number of tasks for image upload. Field introduced in 21.1.3. Allowed in enterprise with any value edition, enterprise with cloud services edition.
* `type` - (Optional) Type of the image patch/system. Enum options - IMAGE_TYPE_PATCH, IMAGE_TYPE_SYSTEM, IMAGE_TYPE_MUST_CHECK. Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.
* `uber_bundle` - (Optional) Status to check if the image is an uber bundle. Field introduced in 18.2.8, 20.1.1. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the image. Field introduced in 18.2.6. Allowed in enterprise with any value edition, essentials edition, basic edition, enterprise with cloud services edition.

