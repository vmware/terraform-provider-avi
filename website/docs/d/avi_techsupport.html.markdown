<!--
    Copyright 2021 VMware, Inc.
    SPDX-License-Identifier: Mozilla Public License 2.0
-->
---
layout: "avi"
page_title: "AVI: avi_techsupport"
sidebar_current: "docs-avi-datasource-techsupport"
description: |-
  Get information of Avi TechSupport.
---

# avi_techsupport

This data source is used to to get avi_techsupport objects.

## Example Usage

```hcl
data "avi_techsupport" "foo_techsupport" {
    uuid = "techsupport-f9cf6b3e-a411-436f-95e2-2982ba2b217b"
    name = "foo"
}
```

## Argument Reference

* `name` - (Optional) Search TechSupport by name.
* `uuid` - (Optional) Search TechSupport by uuid.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `case_number` - 'customer case number for which this tech-upport is generated. ''useful for connected portal and other use-cases.'. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `description` - User provided description to capture additional details and context regarding the tech-support invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `duration` - Total time taken for tech-support collection. Field introduced in 31.2.1. Unit is sec. Allowed with any value in enterprise, enterprise with cloud services edition.
* `end_time` - End timestamp of tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `errors` - Error logged during tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `level` - Name of the tech-support level. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `name` - Name of tech-support invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `node` - Cluster member node on which the tech-support tarball bundle is saved. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_name` - Object name if one exists. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `obj_uuid` - Tech-support collection object uuid specified for different objects such as se/vs/pool etc. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `output` - Tech-support collection output file path. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `params` - Tech-support params associated with latest tech-support collection.user passed params will have more preference. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `progress` - Tech-support collection progress which holds value between 0-100. Allowed values are 0-100. Field introduced in 31.2.1. Unit is percent. Allowed with any value in enterprise, enterprise with cloud services edition.
* `size` - Size of collected tech-support tarball. Field introduced in 31.2.1. Unit is mb. Allowed with any value in enterprise, enterprise with cloud services edition.
* `start_time` - Start timestamp of tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `state` - State of current/last tech-support invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tasks` - Events performed for tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tasks_completed` - Completed set of tasks in the tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `techsupport_readiness` - Techsupport readiness checks execution details. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `tenant_ref` - Tenant uuid associated with the tech-support. It is a reference to an object of type tenant. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `total_tasks` - Total number of tasks in the tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `uuid` - Uuid identifier for the tech-support invocation. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.
* `warnings` - Warning logged during tech-support collection. Field introduced in 31.2.1. Allowed with any value in enterprise, enterprise with cloud services edition.

