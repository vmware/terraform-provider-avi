<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_apispecgenerate"
sidebar_current: "docs-avi-datasource-apispecgenerate"
description: |-
  Get information of Avi ApiSpecGenerate.
---

# avi_apispecgenerate

This data source is used to to get avi_apispecgenerate objects.

## Example Usage

```hcl
data "avi_apispecgenerate" "foo_apispecgenerate" {
    uuid = "apispecgenerate-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search ApiSpecGenerate by name.
* `uuid` - (Optional) Search ApiSpecGenerate by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `completed_events` - Number of tasks completed. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `duration` - Spec generation duration in seconds. Field introduced in 32.1.4. Unit is sec. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `end_time` - Time the spec generation completed or failed. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `name` - Name of the spec generation object. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `params` - Parameters for the spec generation. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `path` - Path to the generated spec file. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `progress` - Overall spec generation progress percentage. Allowed values are 0-100. Field introduced in 32.1.4. Unit is percent. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `start_time` - Time the spec generation started. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `state` - Current lifecycle state of the spec generation. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `task_events` - Per-task status and event details. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `tenant_ref` - It is a reference to an object of type tenant. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `total_events` - Total number of tasks. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.
* `uuid` - Uuid of the spec generation object. Field introduced in 32.1.4. Allowed with any value in enterprise, essentials, basic, enterprise with cloud services edition.

