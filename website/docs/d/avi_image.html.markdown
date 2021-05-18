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

* `controller_info` - Controller package details.
* `controller_patch_name` - Mandatory controller patch name that is applied along with this base image.
* `controller_patch_uuid` - It references the controller-patch associated with the uber image.
* `migrations` - This field describes the api migration related information.
* `name` - Name of the image.
* `se_info` - Se package details.
* `se_patch_name` - Mandatory serviceengine patch name that is applied along with this base image.
* `se_patch_uuid` - It references the service engine patch associated with the uber image.
* `status` - Status to check if the image is present.
* `tenant_ref` - Tenant that this object belongs to.
* `type` - Type of the image patch/system.
* `uber_bundle` - Status to check if the image is an uber bundle.
* `uuid` - Uuid of the image.

