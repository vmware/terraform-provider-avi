############################################################################
# ------------------------------------------------------------------------
# Copyright 2020 VMware, Inc.  All rights reserved. VMware Confidential
# ------------------------------------------------------------------------
###

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

* `controller_info` - (Optional) Controller package details.
* `controller_patch_name` - (Optional) Mandatory controller patch name that is applied along with this base image.
* `controller_patch_uuid` - (Optional) It references the controller-patch associated with the uber image.
* `migrations` - (Optional) This field describes the api migration related information.
* `name` - (Optional) Name of the image.
* `se_info` - (Optional) Se package details.
* `se_patch_name` - (Optional) Mandatory serviceengine patch name that is applied along with this base image.
* `se_patch_uuid` - (Optional) It references the service engine patch associated with the uber image.
* `status` - (Optional) Status to check if the image is present.
* `tenant_ref` - (Optional) Tenant that this object belongs to.
* `type` - (Optional) Type of the image patch/system.
* `uber_bundle` - (Optional) Status to check if the image is an uber bundle.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid of the image.

