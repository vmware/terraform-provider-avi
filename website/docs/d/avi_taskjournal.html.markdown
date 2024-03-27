<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_taskjournal"
sidebar_current: "docs-avi-datasource-taskjournal"
description: |-
  Get information of Avi TaskJournal.
---

# avi_taskjournal

This data source is used to to get avi_taskjournal objects.

## Example Usage

```hcl
data "avi_taskjournal" "foo_taskjournal" {
    uuid = "taskjournal-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TaskJournal by name.
* `uuid` - (Optional) Search TaskJournal by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `errors` - List of errors in the process. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `image_ref` - Image uuid for identifying the current base image. It is a reference to an object of type image. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `info` - Detailed information of journal. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `name` - Name for the task journal. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `obj_cloud_ref` - Cloud that this object belongs to. It is a reference to an object of type cloud. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `operation` - Operation for which the task journal created. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `patch_image_ref` - Image uuid for identifying the current patch. It is a reference to an object of type image. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `summary` - Summary of journal. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `tenant_ref` - Tenant uuid associated with the object. It is a reference to an object of type tenant. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.
* `uuid` - Uuid identifier for the task journal. Field introduced in 30.2.1. Allowed in enterprise edition with any value, enterprise with cloud services edition.

