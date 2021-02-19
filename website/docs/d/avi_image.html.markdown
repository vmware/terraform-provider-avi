############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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
* `migrations` - This field describes the api migration related information. Field introduced in 18.2.6.
* `name` - Name of the image. Field introduced in 18.2.6.
* `se_info` - Se package details. Field introduced in 18.2.6.
* `se_patch_name` - Mandatory serviceengine patch name that is applied along with this base image. Field introduced in 18.2.10, 20.1.1.
* `se_patch_uuid` - It references the service engine patch associated with the uber image. Field introduced in 18.2.8, 20.1.1.
* `status` - Status to check if the image is present. Enum options - SYSERR_SUCCESS, SYSERR_FAILURE, SYSERR_OUT_OF_MEMORY, SYSERR_NO_ENT, SYSERR_INVAL, SYSERR_ACCESS, SYSERR_FAULT, SYSERR_IO, SYSERR_TIMEOUT, SYSERR_NOT_SUPPORTED, SYSERR_NOT_READY, SYSERR_UPGRADE_IN_PROGRESS, SYSERR_WARM_START_IN_PROGRESS, SYSERR_TRY_AGAIN, SYSERR_NOT_UPGRADING, SYSERR_PENDING, SYSERR_EVENT_GEN_FAILURE, SYSERR_CONFIG_PARAM_MISSING, SYSERR_RANGE, SYSERR_BAD_REQUEST... Field introduced in 18.2.6.
* `tenant_ref` - Tenant that this object belongs to. It is a reference to an object of type tenant. Field introduced in 18.2.6.
* `type` - Type of the image patch/system. Enum options - IMAGE_TYPE_PATCH, IMAGE_TYPE_SYSTEM, IMAGE_TYPE_MUST_CHECK. Field introduced in 18.2.6.
* `uber_bundle` - Status to check if the image is an uber bundle. Field introduced in 18.2.8, 20.1.1.
* `uuid` - Uuid of the image. Field introduced in 18.2.6.

