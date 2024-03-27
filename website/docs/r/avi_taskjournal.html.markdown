<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "Avi: avi_taskjournal"
sidebar_current: "docs-avi-resource-taskjournal"
description: |-
  Creates and manages Avi TaskJournal.
---

# avi_taskjournal

The TaskJournal resource allows the creation and management of Avi TaskJournal

## Example Usage

```hcl
resource "avi_taskjournal" "foo" {
    name = "terraform-example-foo"
    tenant_ref = "/api/tenant/?name=admin"
}
```

## Argument Reference

The following arguments are supported:

* `summary` - (Required) Summary of journal. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `errors` - (Optional) List of errors in the process. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `image_ref` - (Optional) Image uuid for identifying the current base image. It is a reference to an object of type image. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `info` - (Optional) Detailed information of journal. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - (Optional) Name for the task journal. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `obj_cloud_ref` - (Optional) Cloud that this object belongs to. It is a reference to an object of type cloud. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `operation` - (Optional) Operation for which the task journal created. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `patch_image_ref` - (Optional) Image uuid for identifying the current patch. It is a reference to an object of type image. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - (Optional) Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.


### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:

* `create` - (Defaults to 40 mins) Used when creating the AMI
* `update` - (Defaults to 40 mins) Used when updating the AMI
* `delete` - (Defaults to 90 mins) Used when deregistering the AMI

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `uuid` -  Uuid identifier for the task journal. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

